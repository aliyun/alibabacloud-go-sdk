// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AbortRunRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// run-13BeYuxc0fxO24uA75UppTiMRoQ
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s AbortRunRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortRunRequest) GoString() string {
	return s.String()
}

func (s *AbortRunRequest) SetRunId(v string) *AbortRunRequest {
	s.RunId = &v
	return s
}

func (s *AbortRunRequest) SetWorkspace(v string) *AbortRunRequest {
	s.Workspace = &v
	return s
}

type AbortRunResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbortRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbortRunResponseBody) GoString() string {
	return s.String()
}

func (s *AbortRunResponseBody) SetHostId(v string) *AbortRunResponseBody {
	s.HostId = &v
	return s
}

func (s *AbortRunResponseBody) SetRequestId(v string) *AbortRunResponseBody {
	s.RequestId = &v
	return s
}

type AbortRunResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbortRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbortRunResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortRunResponse) GoString() string {
	return s.String()
}

func (s *AbortRunResponse) SetHeaders(v map[string]*string) *AbortRunResponse {
	s.Headers = v
	return s
}

func (s *AbortRunResponse) SetStatusCode(v int32) *AbortRunResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortRunResponse) SetBody(v *AbortRunResponseBody) *AbortRunResponse {
	s.Body = v
	return s
}

type AbortSubmissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sub-13BeXWnGvrOl3MC124tazcJVKaV
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s AbortSubmissionRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortSubmissionRequest) GoString() string {
	return s.String()
}

func (s *AbortSubmissionRequest) SetSubmissionId(v string) *AbortSubmissionRequest {
	s.SubmissionId = &v
	return s
}

func (s *AbortSubmissionRequest) SetWorkspace(v string) *AbortSubmissionRequest {
	s.Workspace = &v
	return s
}

type AbortSubmissionResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbortSubmissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbortSubmissionResponseBody) GoString() string {
	return s.String()
}

func (s *AbortSubmissionResponseBody) SetHostId(v string) *AbortSubmissionResponseBody {
	s.HostId = &v
	return s
}

func (s *AbortSubmissionResponseBody) SetRequestId(v string) *AbortSubmissionResponseBody {
	s.RequestId = &v
	return s
}

type AbortSubmissionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbortSubmissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbortSubmissionResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortSubmissionResponse) GoString() string {
	return s.String()
}

func (s *AbortSubmissionResponse) SetHeaders(v map[string]*string) *AbortSubmissionResponse {
	s.Headers = v
	return s
}

func (s *AbortSubmissionResponse) SetStatusCode(v int32) *AbortSubmissionResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortSubmissionResponse) SetBody(v *AbortSubmissionResponseBody) *AbortSubmissionResponse {
	s.Body = v
	return s
}

type CopyPublicEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TestDataset
	Dataset *string `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CopyPublicEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyPublicEntityRequest) GoString() string {
	return s.String()
}

func (s *CopyPublicEntityRequest) SetDataset(v string) *CopyPublicEntityRequest {
	s.Dataset = &v
	return s
}

func (s *CopyPublicEntityRequest) SetEntityType(v string) *CopyPublicEntityRequest {
	s.EntityType = &v
	return s
}

func (s *CopyPublicEntityRequest) SetWorkspace(v string) *CopyPublicEntityRequest {
	s.Workspace = &v
	return s
}

type CopyPublicEntityResponseBody struct {
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CopyPublicEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyPublicEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CopyPublicEntityResponseBody) SetEntityType(v string) *CopyPublicEntityResponseBody {
	s.EntityType = &v
	return s
}

func (s *CopyPublicEntityResponseBody) SetHostId(v string) *CopyPublicEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *CopyPublicEntityResponseBody) SetRequestId(v string) *CopyPublicEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyPublicEntityResponseBody) SetWorkspace(v string) *CopyPublicEntityResponseBody {
	s.Workspace = &v
	return s
}

type CopyPublicEntityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyPublicEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyPublicEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyPublicEntityResponse) GoString() string {
	return s.String()
}

func (s *CopyPublicEntityResponse) SetHeaders(v map[string]*string) *CopyPublicEntityResponse {
	s.Headers = v
	return s
}

func (s *CopyPublicEntityResponse) SetStatusCode(v int32) *CopyPublicEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyPublicEntityResponse) SetBody(v *CopyPublicEntityResponseBody) *CopyPublicEntityResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// App
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// TestToken
	ClientToken *string                    `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Configs     []*CreateAppRequestConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// import "./task/echo.wdl" as echo workflow wf_echo {   call echo.echo }
	Definition   *string                         `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Dependencies []*CreateAppRequestDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
	// example:
	//
	// This is a test app
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Example documentation content
	Documentation *string `json:"Documentation,omitempty" xml:"Documentation,omitempty"`
	// example:
	//
	// {"key":"value"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// WDL
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 1.0
	LanguageVersion *string `json:"LanguageVersion,omitempty" xml:"LanguageVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// main.wdl
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// Example revision comment
	RevisionComment *string `json:"RevisionComment,omitempty" xml:"RevisionComment,omitempty"`
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetAppType(v string) *CreateAppRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppRequest) SetClientToken(v string) *CreateAppRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppRequest) SetConfigs(v []*CreateAppRequestConfigs) *CreateAppRequest {
	s.Configs = v
	return s
}

func (s *CreateAppRequest) SetDefinition(v string) *CreateAppRequest {
	s.Definition = &v
	return s
}

func (s *CreateAppRequest) SetDependencies(v []*CreateAppRequestDependencies) *CreateAppRequest {
	s.Dependencies = v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetDocumentation(v string) *CreateAppRequest {
	s.Documentation = &v
	return s
}

func (s *CreateAppRequest) SetLabels(v string) *CreateAppRequest {
	s.Labels = &v
	return s
}

func (s *CreateAppRequest) SetLanguage(v string) *CreateAppRequest {
	s.Language = &v
	return s
}

func (s *CreateAppRequest) SetLanguageVersion(v string) *CreateAppRequest {
	s.LanguageVersion = &v
	return s
}

func (s *CreateAppRequest) SetPath(v string) *CreateAppRequest {
	s.Path = &v
	return s
}

func (s *CreateAppRequest) SetRevisionComment(v string) *CreateAppRequest {
	s.RevisionComment = &v
	return s
}

func (s *CreateAppRequest) SetRevisionTag(v string) *CreateAppRequest {
	s.RevisionTag = &v
	return s
}

func (s *CreateAppRequest) SetWorkspace(v string) *CreateAppRequest {
	s.Workspace = &v
	return s
}

type CreateAppRequestConfigs struct {
	// example:
	//
	// {     "wgs.apply_bqsr.cpu": 4,     "wgs.apply_bqsr.disks": "local-disk 250 cloud_ssd",     "wgs.apply_bqsr.gatk_Launcher": "/usr/local/bin/gatk-4.1.4.1/gatk",     "wgs.apply_bqsr.java_opts": "\\"-XX:GCTimeLimit=50 -XX:GCHeapFreeLimit=10 -Xms3000m -Djava.io.tmpdir=/mnt\\"",     "wgs.apply_bqsr.memory": "8G"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// ./config/example.json
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateAppRequestConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestConfigs) GoString() string {
	return s.String()
}

func (s *CreateAppRequestConfigs) SetContent(v string) *CreateAppRequestConfigs {
	s.Content = &v
	return s
}

func (s *CreateAppRequestConfigs) SetPath(v string) *CreateAppRequestConfigs {
	s.Path = &v
	return s
}

type CreateAppRequestDependencies struct {
	// example:
	//
	// task echo {   String txt_out = "default_txt_out"   String dd_out = "default_dd_out"    command {     echo Hello > ${txt_out}     /bin/bash -c "dd if=/dev/zero of=${dd_out} bs=1M count=15"     echo "This is a test std out."   }    runtime {     cpu: "2"     memory: "4G"     autoReleaseJob: false     reserveOnFail: true     userData: "key1 value1"     disks: "local-disk 40 cloud_ssd, /home/mount/ 40 cloud_efficiency"     tag: "TestTag"   }    output {     File outTxtFile = "${txt_out}"     File outDdFile = "${dd_out}"   } }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// ./task/echo.wdl
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateAppRequestDependencies) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestDependencies) GoString() string {
	return s.String()
}

func (s *CreateAppRequestDependencies) SetContent(v string) *CreateAppRequestDependencies {
	s.Content = &v
	return s
}

func (s *CreateAppRequestDependencies) SetPath(v string) *CreateAppRequestDependencies {
	s.Path = &v
	return s
}

type CreateAppShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// App
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// TestToken
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigsShrink *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// import "./task/echo.wdl" as echo workflow wf_echo {   call echo.echo }
	Definition         *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	DependenciesShrink *string `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	// example:
	//
	// This is a test app
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Example documentation content
	Documentation *string `json:"Documentation,omitempty" xml:"Documentation,omitempty"`
	// example:
	//
	// {"key":"value"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// WDL
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 1.0
	LanguageVersion *string `json:"LanguageVersion,omitempty" xml:"LanguageVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// main.wdl
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// Example revision comment
	RevisionComment *string `json:"RevisionComment,omitempty" xml:"RevisionComment,omitempty"`
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateAppShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppShrinkRequest) SetAppName(v string) *CreateAppShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppShrinkRequest) SetAppType(v string) *CreateAppShrinkRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppShrinkRequest) SetClientToken(v string) *CreateAppShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppShrinkRequest) SetConfigsShrink(v string) *CreateAppShrinkRequest {
	s.ConfigsShrink = &v
	return s
}

func (s *CreateAppShrinkRequest) SetDefinition(v string) *CreateAppShrinkRequest {
	s.Definition = &v
	return s
}

func (s *CreateAppShrinkRequest) SetDependenciesShrink(v string) *CreateAppShrinkRequest {
	s.DependenciesShrink = &v
	return s
}

func (s *CreateAppShrinkRequest) SetDescription(v string) *CreateAppShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAppShrinkRequest) SetDocumentation(v string) *CreateAppShrinkRequest {
	s.Documentation = &v
	return s
}

func (s *CreateAppShrinkRequest) SetLabels(v string) *CreateAppShrinkRequest {
	s.Labels = &v
	return s
}

func (s *CreateAppShrinkRequest) SetLanguage(v string) *CreateAppShrinkRequest {
	s.Language = &v
	return s
}

func (s *CreateAppShrinkRequest) SetLanguageVersion(v string) *CreateAppShrinkRequest {
	s.LanguageVersion = &v
	return s
}

func (s *CreateAppShrinkRequest) SetPath(v string) *CreateAppShrinkRequest {
	s.Path = &v
	return s
}

func (s *CreateAppShrinkRequest) SetRevisionComment(v string) *CreateAppShrinkRequest {
	s.RevisionComment = &v
	return s
}

func (s *CreateAppShrinkRequest) SetRevisionTag(v string) *CreateAppShrinkRequest {
	s.RevisionTag = &v
	return s
}

func (s *CreateAppShrinkRequest) SetWorkspace(v string) *CreateAppShrinkRequest {
	s.Workspace = &v
	return s
}

type CreateAppResponseBody struct {
	// example:
	//
	// TestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetAppName(v string) *CreateAppResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAppResponseBody) SetHostId(v string) *CreateAppResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetRevision(v string) *CreateAppResponseBody {
	s.Revision = &v
	return s
}

func (s *CreateAppResponseBody) SetRevisionTag(v string) *CreateAppResponseBody {
	s.RevisionTag = &v
	return s
}

func (s *CreateAppResponseBody) SetWorkspace(v string) *CreateAppResponseBody {
	s.Workspace = &v
	return s
}

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetStatusCode(v int32) *CreateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateEntityRequest struct {
	// example:
	//
	// T93cy1mMDE4OTgxNGU1ZTkK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	EntityItems []*CreateEntityRequestEntityItems `json:"EntityItems,omitempty" xml:"EntityItems,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateEntityRequest) SetClientToken(v string) *CreateEntityRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEntityRequest) SetEntityItems(v []*CreateEntityRequestEntityItems) *CreateEntityRequest {
	s.EntityItems = v
	return s
}

func (s *CreateEntityRequest) SetEntityType(v string) *CreateEntityRequest {
	s.EntityType = &v
	return s
}

func (s *CreateEntityRequest) SetWorkspace(v string) *CreateEntityRequest {
	s.Workspace = &v
	return s
}

type CreateEntityRequestEntityItems struct {
	EntityData map[string]*string `json:"EntityData,omitempty" xml:"EntityData,omitempty"`
	// example:
	//
	// s1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
}

func (s CreateEntityRequestEntityItems) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityRequestEntityItems) GoString() string {
	return s.String()
}

func (s *CreateEntityRequestEntityItems) SetEntityData(v map[string]*string) *CreateEntityRequestEntityItems {
	s.EntityData = v
	return s
}

func (s *CreateEntityRequestEntityItems) SetEntityName(v string) *CreateEntityRequestEntityItems {
	s.EntityName = &v
	return s
}

type CreateEntityShrinkRequest struct {
	// example:
	//
	// T93cy1mMDE4OTgxNGU1ZTkK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	EntityItemsShrink *string `json:"EntityItems,omitempty" xml:"EntityItems,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateEntityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEntityShrinkRequest) SetClientToken(v string) *CreateEntityShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEntityShrinkRequest) SetEntityItemsShrink(v string) *CreateEntityShrinkRequest {
	s.EntityItemsShrink = &v
	return s
}

func (s *CreateEntityShrinkRequest) SetEntityType(v string) *CreateEntityShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *CreateEntityShrinkRequest) SetWorkspace(v string) *CreateEntityShrinkRequest {
	s.Workspace = &v
	return s
}

type CreateEntityResponseBody struct {
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEntityResponseBody) SetEntityType(v string) *CreateEntityResponseBody {
	s.EntityType = &v
	return s
}

func (s *CreateEntityResponseBody) SetHostId(v string) *CreateEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateEntityResponseBody) SetRequestId(v string) *CreateEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEntityResponseBody) SetWorkspace(v string) *CreateEntityResponseBody {
	s.Workspace = &v
	return s
}

type CreateEntityResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateEntityResponse) SetHeaders(v map[string]*string) *CreateEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateEntityResponse) SetStatusCode(v int32) *CreateEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEntityResponse) SetBody(v *CreateEntityResponseBody) *CreateEntityResponse {
	s.Body = v
	return s
}

type CreateRunRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Sentieon
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 16
	AppRevision *string `json:"AppRevision,omitempty" xml:"AppRevision,omitempty"`
	// example:
	//
	// MTIzNDU2P1dvcmtzcGFjZT93cy1mMDE4OTgxNGU1ZTkK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {}
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	// example:
	//
	// test run
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// oss://my-bucket/exec/
	ExecuteDirectory *string `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	// example:
	//
	// { 	"CallCaching": true,     "FailureMode": "NoNewCalls" }
	ExecuteOptions *CreateRunRequestExecuteOptions `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// {         "workflow1.task1.param1": "helloooo.txt",         "workflow1.task1.param2": "abc",         "workflow1.task2.param1": "new_helloooo.txt",         "workflow1.task2.param2": "def"     }
	Inputs *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// example:
	//
	// {"env": "test"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// oss://my-bucket/output/
	OutputFolder *string `json:"OutputFolder,omitempty" xml:"OutputFolder,omitempty"`
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// example:
	//
	// TestRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// wgs_sentieon_hc_with_qc_NGPTS1900029301
	RunName *string `json:"RunName,omitempty" xml:"RunName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateRunRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRunRequest) GoString() string {
	return s.String()
}

func (s *CreateRunRequest) SetAppName(v string) *CreateRunRequest {
	s.AppName = &v
	return s
}

func (s *CreateRunRequest) SetAppRevision(v string) *CreateRunRequest {
	s.AppRevision = &v
	return s
}

func (s *CreateRunRequest) SetClientToken(v string) *CreateRunRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRunRequest) SetDefaultRuntime(v string) *CreateRunRequest {
	s.DefaultRuntime = &v
	return s
}

func (s *CreateRunRequest) SetDescription(v string) *CreateRunRequest {
	s.Description = &v
	return s
}

func (s *CreateRunRequest) SetExecuteDirectory(v string) *CreateRunRequest {
	s.ExecuteDirectory = &v
	return s
}

func (s *CreateRunRequest) SetExecuteOptions(v *CreateRunRequestExecuteOptions) *CreateRunRequest {
	s.ExecuteOptions = v
	return s
}

func (s *CreateRunRequest) SetInputs(v string) *CreateRunRequest {
	s.Inputs = &v
	return s
}

func (s *CreateRunRequest) SetLabels(v string) *CreateRunRequest {
	s.Labels = &v
	return s
}

func (s *CreateRunRequest) SetOutputFolder(v string) *CreateRunRequest {
	s.OutputFolder = &v
	return s
}

func (s *CreateRunRequest) SetRevisionTag(v string) *CreateRunRequest {
	s.RevisionTag = &v
	return s
}

func (s *CreateRunRequest) SetRole(v string) *CreateRunRequest {
	s.Role = &v
	return s
}

func (s *CreateRunRequest) SetRunName(v string) *CreateRunRequest {
	s.RunName = &v
	return s
}

func (s *CreateRunRequest) SetWorkspace(v string) *CreateRunRequest {
	s.Workspace = &v
	return s
}

type CreateRunRequestExecuteOptions struct {
	// example:
	//
	// True
	CallCaching *bool `json:"CallCaching,omitempty" xml:"CallCaching,omitempty"`
	// example:
	//
	// Mount
	DataLoader *string `json:"DataLoader,omitempty" xml:"DataLoader,omitempty"`
	// example:
	//
	// True
	DeleteIntermediateResults *bool `json:"DeleteIntermediateResults,omitempty" xml:"DeleteIntermediateResults,omitempty"`
	// example:
	//
	// NoNewCalls
	FailureMode *string `json:"FailureMode,omitempty" xml:"FailureMode,omitempty"`
	// example:
	//
	// True
	UseRelativeOutputPaths *bool `json:"UseRelativeOutputPaths,omitempty" xml:"UseRelativeOutputPaths,omitempty"`
}

func (s CreateRunRequestExecuteOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateRunRequestExecuteOptions) GoString() string {
	return s.String()
}

func (s *CreateRunRequestExecuteOptions) SetCallCaching(v bool) *CreateRunRequestExecuteOptions {
	s.CallCaching = &v
	return s
}

func (s *CreateRunRequestExecuteOptions) SetDataLoader(v string) *CreateRunRequestExecuteOptions {
	s.DataLoader = &v
	return s
}

func (s *CreateRunRequestExecuteOptions) SetDeleteIntermediateResults(v bool) *CreateRunRequestExecuteOptions {
	s.DeleteIntermediateResults = &v
	return s
}

func (s *CreateRunRequestExecuteOptions) SetFailureMode(v string) *CreateRunRequestExecuteOptions {
	s.FailureMode = &v
	return s
}

func (s *CreateRunRequestExecuteOptions) SetUseRelativeOutputPaths(v bool) *CreateRunRequestExecuteOptions {
	s.UseRelativeOutputPaths = &v
	return s
}

type CreateRunShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Sentieon
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 16
	AppRevision *string `json:"AppRevision,omitempty" xml:"AppRevision,omitempty"`
	// example:
	//
	// MTIzNDU2P1dvcmtzcGFjZT93cy1mMDE4OTgxNGU1ZTkK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {}
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	// example:
	//
	// test run
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// oss://my-bucket/exec/
	ExecuteDirectory *string `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	// example:
	//
	// { 	"CallCaching": true,     "FailureMode": "NoNewCalls" }
	ExecuteOptionsShrink *string `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {         "workflow1.task1.param1": "helloooo.txt",         "workflow1.task1.param2": "abc",         "workflow1.task2.param1": "new_helloooo.txt",         "workflow1.task2.param2": "def"     }
	Inputs *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// example:
	//
	// {"env": "test"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// oss://my-bucket/output/
	OutputFolder *string `json:"OutputFolder,omitempty" xml:"OutputFolder,omitempty"`
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// example:
	//
	// TestRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// wgs_sentieon_hc_with_qc_NGPTS1900029301
	RunName *string `json:"RunName,omitempty" xml:"RunName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateRunShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRunShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRunShrinkRequest) SetAppName(v string) *CreateRunShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateRunShrinkRequest) SetAppRevision(v string) *CreateRunShrinkRequest {
	s.AppRevision = &v
	return s
}

func (s *CreateRunShrinkRequest) SetClientToken(v string) *CreateRunShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRunShrinkRequest) SetDefaultRuntime(v string) *CreateRunShrinkRequest {
	s.DefaultRuntime = &v
	return s
}

func (s *CreateRunShrinkRequest) SetDescription(v string) *CreateRunShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRunShrinkRequest) SetExecuteDirectory(v string) *CreateRunShrinkRequest {
	s.ExecuteDirectory = &v
	return s
}

func (s *CreateRunShrinkRequest) SetExecuteOptionsShrink(v string) *CreateRunShrinkRequest {
	s.ExecuteOptionsShrink = &v
	return s
}

func (s *CreateRunShrinkRequest) SetInputs(v string) *CreateRunShrinkRequest {
	s.Inputs = &v
	return s
}

func (s *CreateRunShrinkRequest) SetLabels(v string) *CreateRunShrinkRequest {
	s.Labels = &v
	return s
}

func (s *CreateRunShrinkRequest) SetOutputFolder(v string) *CreateRunShrinkRequest {
	s.OutputFolder = &v
	return s
}

func (s *CreateRunShrinkRequest) SetRevisionTag(v string) *CreateRunShrinkRequest {
	s.RevisionTag = &v
	return s
}

func (s *CreateRunShrinkRequest) SetRole(v string) *CreateRunShrinkRequest {
	s.Role = &v
	return s
}

func (s *CreateRunShrinkRequest) SetRunName(v string) *CreateRunShrinkRequest {
	s.RunName = &v
	return s
}

func (s *CreateRunShrinkRequest) SetWorkspace(v string) *CreateRunShrinkRequest {
	s.Workspace = &v
	return s
}

type CreateRunResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// run-13BeYuxc0fxO24uA75UppTiMRoQ
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBody) SetHostId(v string) *CreateRunResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateRunResponseBody) SetRequestId(v string) *CreateRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRunResponseBody) SetRunId(v string) *CreateRunResponseBody {
	s.RunId = &v
	return s
}

func (s *CreateRunResponseBody) SetWorkspace(v string) *CreateRunResponseBody {
	s.Workspace = &v
	return s
}

type CreateRunResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRunResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRunResponse) GoString() string {
	return s.String()
}

func (s *CreateRunResponse) SetHeaders(v map[string]*string) *CreateRunResponse {
	s.Headers = v
	return s
}

func (s *CreateRunResponse) SetStatusCode(v int32) *CreateRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRunResponse) SetBody(v *CreateRunResponseBody) *CreateRunResponse {
	s.Body = v
	return s
}

type CreateSubmissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// TestToken
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {}
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	// if can be null:
	// false
	EntityNames []*string `json:"EntityNames,omitempty" xml:"EntityNames,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// TestEntityType
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// oss://my-bucket/my-exe-dir/
	ExecuteDirectory *string `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	// example:
	//
	// { 	        "CallCaching": false, 	        "DeleteIntermediateResults": true, 	        "FailureMode": "NoNewCalls"         }
	ExecuteOptions *string `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty"`
	// example:
	//
	// {}
	Inputs *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// example:
	//
	// oss://my-bucket/my-output-dir/
	OutputFolder *string `json:"OutputFolder,omitempty" xml:"OutputFolder,omitempty"`
	// example:
	//
	// {}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// example:
	//
	// 1
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateSubmissionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubmissionRequest) GoString() string {
	return s.String()
}

func (s *CreateSubmissionRequest) SetAppName(v string) *CreateSubmissionRequest {
	s.AppName = &v
	return s
}

func (s *CreateSubmissionRequest) SetClientToken(v string) *CreateSubmissionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSubmissionRequest) SetDefaultRuntime(v string) *CreateSubmissionRequest {
	s.DefaultRuntime = &v
	return s
}

func (s *CreateSubmissionRequest) SetEntityNames(v []*string) *CreateSubmissionRequest {
	s.EntityNames = v
	return s
}

func (s *CreateSubmissionRequest) SetEntityType(v string) *CreateSubmissionRequest {
	s.EntityType = &v
	return s
}

func (s *CreateSubmissionRequest) SetExecuteDirectory(v string) *CreateSubmissionRequest {
	s.ExecuteDirectory = &v
	return s
}

func (s *CreateSubmissionRequest) SetExecuteOptions(v string) *CreateSubmissionRequest {
	s.ExecuteOptions = &v
	return s
}

func (s *CreateSubmissionRequest) SetInputs(v string) *CreateSubmissionRequest {
	s.Inputs = &v
	return s
}

func (s *CreateSubmissionRequest) SetOutputFolder(v string) *CreateSubmissionRequest {
	s.OutputFolder = &v
	return s
}

func (s *CreateSubmissionRequest) SetOutputs(v string) *CreateSubmissionRequest {
	s.Outputs = &v
	return s
}

func (s *CreateSubmissionRequest) SetRevision(v string) *CreateSubmissionRequest {
	s.Revision = &v
	return s
}

func (s *CreateSubmissionRequest) SetRevisionTag(v string) *CreateSubmissionRequest {
	s.RevisionTag = &v
	return s
}

func (s *CreateSubmissionRequest) SetWorkspace(v string) *CreateSubmissionRequest {
	s.Workspace = &v
	return s
}

type CreateSubmissionShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// TestToken
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {}
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	// if can be null:
	// false
	EntityNamesShrink *string `json:"EntityNames,omitempty" xml:"EntityNames,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestEntityType
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// oss://my-bucket/my-exe-dir/
	ExecuteDirectory *string `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	// example:
	//
	// { 	        "CallCaching": false, 	        "DeleteIntermediateResults": true, 	        "FailureMode": "NoNewCalls"         }
	ExecuteOptions *string `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty"`
	// example:
	//
	// {}
	Inputs *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// example:
	//
	// oss://my-bucket/my-output-dir/
	OutputFolder *string `json:"OutputFolder,omitempty" xml:"OutputFolder,omitempty"`
	// example:
	//
	// {}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// example:
	//
	// 1
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateSubmissionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubmissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSubmissionShrinkRequest) SetAppName(v string) *CreateSubmissionShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetClientToken(v string) *CreateSubmissionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetDefaultRuntime(v string) *CreateSubmissionShrinkRequest {
	s.DefaultRuntime = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetEntityNamesShrink(v string) *CreateSubmissionShrinkRequest {
	s.EntityNamesShrink = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetEntityType(v string) *CreateSubmissionShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetExecuteDirectory(v string) *CreateSubmissionShrinkRequest {
	s.ExecuteDirectory = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetExecuteOptions(v string) *CreateSubmissionShrinkRequest {
	s.ExecuteOptions = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetInputs(v string) *CreateSubmissionShrinkRequest {
	s.Inputs = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetOutputFolder(v string) *CreateSubmissionShrinkRequest {
	s.OutputFolder = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetOutputs(v string) *CreateSubmissionShrinkRequest {
	s.Outputs = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetRevision(v string) *CreateSubmissionShrinkRequest {
	s.Revision = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetRevisionTag(v string) *CreateSubmissionShrinkRequest {
	s.RevisionTag = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetWorkspace(v string) *CreateSubmissionShrinkRequest {
	s.Workspace = &v
	return s
}

type CreateSubmissionResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// sub-13BeXWnGvrOl3MC124tazcJVKaV
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateSubmissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubmissionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubmissionResponseBody) SetHostId(v string) *CreateSubmissionResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateSubmissionResponseBody) SetRequestId(v string) *CreateSubmissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubmissionResponseBody) SetSubmissionId(v string) *CreateSubmissionResponseBody {
	s.SubmissionId = &v
	return s
}

func (s *CreateSubmissionResponseBody) SetWorkspace(v string) *CreateSubmissionResponseBody {
	s.Workspace = &v
	return s
}

type CreateSubmissionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubmissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubmissionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubmissionResponse) GoString() string {
	return s.String()
}

func (s *CreateSubmissionResponse) SetHeaders(v map[string]*string) *CreateSubmissionResponse {
	s.Headers = v
	return s
}

func (s *CreateSubmissionResponse) SetStatusCode(v int32) *CreateSubmissionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubmissionResponse) SetBody(v *CreateSubmissionResponseBody) *CreateSubmissionResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// MyApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 4.1.4.1-v1.0
	AppRevision *string `json:"AppRevision,omitempty" xml:"AppRevision,omitempty"`
	// example:
	//
	// GFjZT93cy1mMDE4OTgxNGU1ZTkK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// This is a test template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InputsExpression []*CreateTemplateRequestInputsExpression `json:"InputsExpression,omitempty" xml:"InputsExpression,omitempty" type:"Repeated"`
	// example:
	//
	// {"key1": "value1"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// This parameter is required.
	OutputsExpression []*CreateTemplateRequestOutputsExpression `json:"OutputsExpression,omitempty" xml:"OutputsExpression,omitempty" type:"Repeated"`
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// example:
	//
	// TestEntity
	RootEntity *string `json:"RootEntity,omitempty" xml:"RootEntity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetAppName(v string) *CreateTemplateRequest {
	s.AppName = &v
	return s
}

func (s *CreateTemplateRequest) SetAppRevision(v string) *CreateTemplateRequest {
	s.AppRevision = &v
	return s
}

func (s *CreateTemplateRequest) SetClientToken(v string) *CreateTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTemplateRequest) SetDescription(v string) *CreateTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequest) SetInputsExpression(v []*CreateTemplateRequestInputsExpression) *CreateTemplateRequest {
	s.InputsExpression = v
	return s
}

func (s *CreateTemplateRequest) SetLabels(v string) *CreateTemplateRequest {
	s.Labels = &v
	return s
}

func (s *CreateTemplateRequest) SetOutputsExpression(v []*CreateTemplateRequestOutputsExpression) *CreateTemplateRequest {
	s.OutputsExpression = v
	return s
}

func (s *CreateTemplateRequest) SetRevisionTag(v string) *CreateTemplateRequest {
	s.RevisionTag = &v
	return s
}

func (s *CreateTemplateRequest) SetRootEntity(v string) *CreateTemplateRequest {
	s.RootEntity = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetWorkspace(v string) *CreateTemplateRequest {
	s.Workspace = &v
	return s
}

type CreateTemplateRequestInputsExpression struct {
	// example:
	//
	// test help
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// 0
	StepOrder *int32 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fastq1
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// File
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/file1.fastq
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s CreateTemplateRequestInputsExpression) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequestInputsExpression) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequestInputsExpression) SetHelp(v string) *CreateTemplateRequestInputsExpression {
	s.Help = &v
	return s
}

func (s *CreateTemplateRequestInputsExpression) SetRequired(v bool) *CreateTemplateRequestInputsExpression {
	s.Required = &v
	return s
}

func (s *CreateTemplateRequestInputsExpression) SetStepOrder(v int32) *CreateTemplateRequestInputsExpression {
	s.StepOrder = &v
	return s
}

func (s *CreateTemplateRequestInputsExpression) SetTaskName(v string) *CreateTemplateRequestInputsExpression {
	s.TaskName = &v
	return s
}

func (s *CreateTemplateRequestInputsExpression) SetVariableName(v string) *CreateTemplateRequestInputsExpression {
	s.VariableName = &v
	return s
}

func (s *CreateTemplateRequestInputsExpression) SetVariableType(v string) *CreateTemplateRequestInputsExpression {
	s.VariableType = &v
	return s
}

func (s *CreateTemplateRequestInputsExpression) SetVariableValue(v string) *CreateTemplateRequestInputsExpression {
	s.VariableValue = &v
	return s
}

type CreateTemplateRequestOutputsExpression struct {
	// example:
	//
	// test help
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// 1
	StepOrder *int32 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bam
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// File
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/file.bam
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s CreateTemplateRequestOutputsExpression) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequestOutputsExpression) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequestOutputsExpression) SetHelp(v string) *CreateTemplateRequestOutputsExpression {
	s.Help = &v
	return s
}

func (s *CreateTemplateRequestOutputsExpression) SetRequired(v bool) *CreateTemplateRequestOutputsExpression {
	s.Required = &v
	return s
}

func (s *CreateTemplateRequestOutputsExpression) SetStepOrder(v int32) *CreateTemplateRequestOutputsExpression {
	s.StepOrder = &v
	return s
}

func (s *CreateTemplateRequestOutputsExpression) SetTaskName(v string) *CreateTemplateRequestOutputsExpression {
	s.TaskName = &v
	return s
}

func (s *CreateTemplateRequestOutputsExpression) SetVariableName(v string) *CreateTemplateRequestOutputsExpression {
	s.VariableName = &v
	return s
}

func (s *CreateTemplateRequestOutputsExpression) SetVariableType(v string) *CreateTemplateRequestOutputsExpression {
	s.VariableType = &v
	return s
}

func (s *CreateTemplateRequestOutputsExpression) SetVariableValue(v string) *CreateTemplateRequestOutputsExpression {
	s.VariableValue = &v
	return s
}

type CreateTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// MyApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 4.1.4.1-v1.0
	AppRevision *string `json:"AppRevision,omitempty" xml:"AppRevision,omitempty"`
	// example:
	//
	// GFjZT93cy1mMDE4OTgxNGU1ZTkK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// This is a test template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InputsExpressionShrink *string `json:"InputsExpression,omitempty" xml:"InputsExpression,omitempty"`
	// example:
	//
	// {"key1": "value1"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// This parameter is required.
	OutputsExpressionShrink *string `json:"OutputsExpression,omitempty" xml:"OutputsExpression,omitempty"`
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// example:
	//
	// TestEntity
	RootEntity *string `json:"RootEntity,omitempty" xml:"RootEntity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateShrinkRequest) SetAppName(v string) *CreateTemplateShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetAppRevision(v string) *CreateTemplateShrinkRequest {
	s.AppRevision = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetClientToken(v string) *CreateTemplateShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetDescription(v string) *CreateTemplateShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetInputsExpressionShrink(v string) *CreateTemplateShrinkRequest {
	s.InputsExpressionShrink = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetLabels(v string) *CreateTemplateShrinkRequest {
	s.Labels = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetOutputsExpressionShrink(v string) *CreateTemplateShrinkRequest {
	s.OutputsExpressionShrink = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetRevisionTag(v string) *CreateTemplateShrinkRequest {
	s.RevisionTag = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetRootEntity(v string) *CreateTemplateShrinkRequest {
	s.RootEntity = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetTemplateName(v string) *CreateTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetWorkspace(v string) *CreateTemplateShrinkRequest {
	s.Workspace = &v
	return s
}

type CreateTemplateResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test-template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetHostId(v string) *CreateTemplateResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplateName(v string) *CreateTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateResponseBody) SetWorkspace(v string) *CreateTemplateResponseBody {
	s.Workspace = &v
	return s
}

type CreateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetStatusCode(v int32) *CreateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type CreateWorkspaceRequest struct {
	// example:
	//
	// TestToken
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// This is a test workspace
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 30
	JobLifecycle *int32 `json:"JobLifecycle,omitempty" xml:"JobLifecycle,omitempty"`
	// example:
	//
	// {"key":"value"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// TestRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://gstor-default-workspace-cn-shanghai-bae3193d/
	Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) SetClientToken(v string) *CreateWorkspaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDescription(v string) *CreateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceRequest) SetJobLifecycle(v int32) *CreateWorkspaceRequest {
	s.JobLifecycle = &v
	return s
}

func (s *CreateWorkspaceRequest) SetLabels(v string) *CreateWorkspaceRequest {
	s.Labels = &v
	return s
}

func (s *CreateWorkspaceRequest) SetRole(v string) *CreateWorkspaceRequest {
	s.Role = &v
	return s
}

func (s *CreateWorkspaceRequest) SetStorage(v string) *CreateWorkspaceRequest {
	s.Storage = &v
	return s
}

func (s *CreateWorkspaceRequest) SetWorkspace(v string) *CreateWorkspaceRequest {
	s.Workspace = &v
	return s
}

type CreateWorkspaceResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) SetHostId(v string) *CreateWorkspaceResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetWorkspace(v string) *CreateWorkspaceResponseBody {
	s.Workspace = &v
	return s
}

type CreateWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponse) SetHeaders(v map[string]*string) *CreateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceResponse) SetStatusCode(v int32) *CreateWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceResponse) SetBody(v *CreateWorkspaceResponseBody) *CreateWorkspaceResponse {
	s.Body = v
	return s
}

type DeleteAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) SetAppName(v string) *DeleteAppRequest {
	s.AppName = &v
	return s
}

func (s *DeleteAppRequest) SetRevision(v string) *DeleteAppRequest {
	s.Revision = &v
	return s
}

func (s *DeleteAppRequest) SetWorkspace(v string) *DeleteAppRequest {
	s.Workspace = &v
	return s
}

type DeleteAppResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetHostId(v string) *DeleteAppResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetStatusCode(v int32) *DeleteAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

type DeleteEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DeleteEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteEntityRequest) SetEntityType(v string) *DeleteEntityRequest {
	s.EntityType = &v
	return s
}

func (s *DeleteEntityRequest) SetWorkspace(v string) *DeleteEntityRequest {
	s.Workspace = &v
	return s
}

type DeleteEntityResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEntityResponseBody) SetHostId(v string) *DeleteEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteEntityResponseBody) SetRequestId(v string) *DeleteEntityResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEntityResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteEntityResponse) SetHeaders(v map[string]*string) *DeleteEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteEntityResponse) SetStatusCode(v int32) *DeleteEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEntityResponse) SetBody(v *DeleteEntityResponseBody) *DeleteEntityResponse {
	s.Body = v
	return s
}

type DeleteEntityItemsRequest struct {
	// This parameter is required.
	EntityNames []*string `json:"EntityNames,omitempty" xml:"EntityNames,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DeleteEntityItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityItemsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEntityItemsRequest) SetEntityNames(v []*string) *DeleteEntityItemsRequest {
	s.EntityNames = v
	return s
}

func (s *DeleteEntityItemsRequest) SetEntityType(v string) *DeleteEntityItemsRequest {
	s.EntityType = &v
	return s
}

func (s *DeleteEntityItemsRequest) SetWorkspace(v string) *DeleteEntityItemsRequest {
	s.Workspace = &v
	return s
}

type DeleteEntityItemsShrinkRequest struct {
	// This parameter is required.
	EntityNamesShrink *string `json:"EntityNames,omitempty" xml:"EntityNames,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DeleteEntityItemsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteEntityItemsShrinkRequest) SetEntityNamesShrink(v string) *DeleteEntityItemsShrinkRequest {
	s.EntityNamesShrink = &v
	return s
}

func (s *DeleteEntityItemsShrinkRequest) SetEntityType(v string) *DeleteEntityItemsShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *DeleteEntityItemsShrinkRequest) SetWorkspace(v string) *DeleteEntityItemsShrinkRequest {
	s.Workspace = &v
	return s
}

type DeleteEntityItemsResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEntityItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEntityItemsResponseBody) SetHostId(v string) *DeleteEntityItemsResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteEntityItemsResponseBody) SetRequestId(v string) *DeleteEntityItemsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEntityItemsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEntityItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEntityItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityItemsResponse) GoString() string {
	return s.String()
}

func (s *DeleteEntityItemsResponse) SetHeaders(v map[string]*string) *DeleteEntityItemsResponse {
	s.Headers = v
	return s
}

func (s *DeleteEntityItemsResponse) SetStatusCode(v int32) *DeleteEntityItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEntityItemsResponse) SetBody(v *DeleteEntityItemsResponseBody) *DeleteEntityItemsResponse {
	s.Body = v
	return s
}

type DeleteRunRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// run-13BeYuxc0fxO24uA75UppTiMRoQ
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DeleteRunRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRunRequest) GoString() string {
	return s.String()
}

func (s *DeleteRunRequest) SetRunId(v string) *DeleteRunRequest {
	s.RunId = &v
	return s
}

func (s *DeleteRunRequest) SetWorkspace(v string) *DeleteRunRequest {
	s.Workspace = &v
	return s
}

type DeleteRunResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRunResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRunResponseBody) SetHostId(v string) *DeleteRunResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteRunResponseBody) SetRequestId(v string) *DeleteRunResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRunResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRunResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRunResponse) GoString() string {
	return s.String()
}

func (s *DeleteRunResponse) SetHeaders(v map[string]*string) *DeleteRunResponse {
	s.Headers = v
	return s
}

func (s *DeleteRunResponse) SetStatusCode(v int32) *DeleteRunResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRunResponse) SetBody(v *DeleteRunResponseBody) *DeleteRunResponse {
	s.Body = v
	return s
}

type DeleteSubmissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sub-13BeXWnGvrOl3MC124tazcJVKaV
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DeleteSubmissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubmissionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubmissionRequest) SetSubmissionId(v string) *DeleteSubmissionRequest {
	s.SubmissionId = &v
	return s
}

func (s *DeleteSubmissionRequest) SetWorkspace(v string) *DeleteSubmissionRequest {
	s.Workspace = &v
	return s
}

type DeleteSubmissionResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSubmissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubmissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubmissionResponseBody) SetHostId(v string) *DeleteSubmissionResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteSubmissionResponseBody) SetRequestId(v string) *DeleteSubmissionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSubmissionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSubmissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSubmissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubmissionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubmissionResponse) SetHeaders(v map[string]*string) *DeleteSubmissionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubmissionResponse) SetStatusCode(v int32) *DeleteSubmissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubmissionResponse) SetBody(v *DeleteSubmissionResponseBody) *DeleteSubmissionResponse {
	s.Body = v
	return s
}

type DeleteTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) SetTemplateName(v string) *DeleteTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *DeleteTemplateRequest) SetWorkspace(v string) *DeleteTemplateRequest {
	s.Workspace = &v
	return s
}

type DeleteTemplateResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) SetHostId(v string) *DeleteTemplateResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetStatusCode(v int32) *DeleteTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
	s.Body = v
	return s
}

type DeleteWorkspaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// MyTestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DeleteWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceRequest) SetWorkspace(v string) *DeleteWorkspaceRequest {
	s.Workspace = &v
	return s
}

type DeleteWorkspaceResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBody) SetHostId(v string) *DeleteWorkspaceResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetRequestId(v string) *DeleteWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceResponse) SetStatusCode(v int32) *DeleteWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceResponse) SetBody(v *DeleteWorkspaceResponseBody) *DeleteWorkspaceResponse {
	s.Body = v
	return s
}

type DownloadEntityRequest struct {
	EntityNames []*string `json:"EntityNames,omitempty" xml:"EntityNames,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DownloadEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadEntityRequest) GoString() string {
	return s.String()
}

func (s *DownloadEntityRequest) SetEntityNames(v []*string) *DownloadEntityRequest {
	s.EntityNames = v
	return s
}

func (s *DownloadEntityRequest) SetEntityType(v string) *DownloadEntityRequest {
	s.EntityType = &v
	return s
}

func (s *DownloadEntityRequest) SetWorkspace(v string) *DownloadEntityRequest {
	s.Workspace = &v
	return s
}

type DownloadEntityShrinkRequest struct {
	EntityNamesShrink *string `json:"EntityNames,omitempty" xml:"EntityNames,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DownloadEntityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *DownloadEntityShrinkRequest) SetEntityNamesShrink(v string) *DownloadEntityShrinkRequest {
	s.EntityNamesShrink = &v
	return s
}

func (s *DownloadEntityShrinkRequest) SetEntityType(v string) *DownloadEntityShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *DownloadEntityShrinkRequest) SetWorkspace(v string) *DownloadEntityShrinkRequest {
	s.Workspace = &v
	return s
}

type DownloadEntityResponseBody struct {
	// example:
	//
	// https://my-bucket.oss.aliyucs.com/entity.csv
	EntityCSVFile *string `json:"EntityCSVFile,omitempty" xml:"EntityCSVFile,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadEntityResponseBody) SetEntityCSVFile(v string) *DownloadEntityResponseBody {
	s.EntityCSVFile = &v
	return s
}

func (s *DownloadEntityResponseBody) SetHostId(v string) *DownloadEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *DownloadEntityResponseBody) SetRequestId(v string) *DownloadEntityResponseBody {
	s.RequestId = &v
	return s
}

type DownloadEntityResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadEntityResponse) GoString() string {
	return s.String()
}

func (s *DownloadEntityResponse) SetHeaders(v map[string]*string) *DownloadEntityResponse {
	s.Headers = v
	return s
}

func (s *DownloadEntityResponse) SetStatusCode(v int32) *DownloadEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadEntityResponse) SetBody(v *DownloadEntityResponseBody) *DownloadEntityResponse {
	s.Body = v
	return s
}

type GetAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 0
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// example:
	//
	// TestTag
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppRequest) GoString() string {
	return s.String()
}

func (s *GetAppRequest) SetAppName(v string) *GetAppRequest {
	s.AppName = &v
	return s
}

func (s *GetAppRequest) SetRevision(v string) *GetAppRequest {
	s.Revision = &v
	return s
}

func (s *GetAppRequest) SetRevisionTag(v string) *GetAppRequest {
	s.RevisionTag = &v
	return s
}

func (s *GetAppRequest) SetWorkspace(v string) *GetAppRequest {
	s.Workspace = &v
	return s
}

type GetAppResponseBody struct {
	// example:
	//
	// MyTestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// App
	AppType *string                      `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Configs []*GetAppResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// WDL content
	Definition   *string                           `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Dependencies []*GetAppResponseBodyDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
	// example:
	//
	// This is a test app
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// example documentation content
	Documentation *string `json:"Documentation,omitempty" xml:"Documentation,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string                     `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Inputs []*GetAppResponseBodyInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	Labels map[string]*string          `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// WDL
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 1.0
	LanguageVersion *string `json:"LanguageVersion,omitempty" xml:"LanguageVersion,omitempty"`
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	LastModifiedTime *string                      `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Outputs          []*GetAppResponseBodyOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// example:
	//
	// main.wdl
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// example:
	//
	// Example revision comment
	RevisionComment *string `json:"RevisionComment,omitempty" xml:"RevisionComment,omitempty"`
	// example:
	//
	// TestTag
	RevisionTag *string                        `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	Revisions   []*GetAppResponseBodyRevisions `json:"Revisions,omitempty" xml:"Revisions,omitempty" type:"Repeated"`
	// example:
	//
	// Public
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// EasyGene
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// Example url
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
	// example:
	//
	// MyTestWorkflow
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppResponseBody) SetAppName(v string) *GetAppResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppResponseBody) SetAppType(v string) *GetAppResponseBody {
	s.AppType = &v
	return s
}

func (s *GetAppResponseBody) SetConfigs(v []*GetAppResponseBodyConfigs) *GetAppResponseBody {
	s.Configs = v
	return s
}

func (s *GetAppResponseBody) SetCreateTime(v string) *GetAppResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetAppResponseBody) SetDefinition(v string) *GetAppResponseBody {
	s.Definition = &v
	return s
}

func (s *GetAppResponseBody) SetDependencies(v []*GetAppResponseBodyDependencies) *GetAppResponseBody {
	s.Dependencies = v
	return s
}

func (s *GetAppResponseBody) SetDescription(v string) *GetAppResponseBody {
	s.Description = &v
	return s
}

func (s *GetAppResponseBody) SetDocumentation(v string) *GetAppResponseBody {
	s.Documentation = &v
	return s
}

func (s *GetAppResponseBody) SetHostId(v string) *GetAppResponseBody {
	s.HostId = &v
	return s
}

func (s *GetAppResponseBody) SetInputs(v []*GetAppResponseBodyInputs) *GetAppResponseBody {
	s.Inputs = v
	return s
}

func (s *GetAppResponseBody) SetLabels(v map[string]*string) *GetAppResponseBody {
	s.Labels = v
	return s
}

func (s *GetAppResponseBody) SetLanguage(v string) *GetAppResponseBody {
	s.Language = &v
	return s
}

func (s *GetAppResponseBody) SetLanguageVersion(v string) *GetAppResponseBody {
	s.LanguageVersion = &v
	return s
}

func (s *GetAppResponseBody) SetLastModifiedTime(v string) *GetAppResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetAppResponseBody) SetOutputs(v []*GetAppResponseBodyOutputs) *GetAppResponseBody {
	s.Outputs = v
	return s
}

func (s *GetAppResponseBody) SetPath(v string) *GetAppResponseBody {
	s.Path = &v
	return s
}

func (s *GetAppResponseBody) SetRequestId(v string) *GetAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppResponseBody) SetRevision(v string) *GetAppResponseBody {
	s.Revision = &v
	return s
}

func (s *GetAppResponseBody) SetRevisionComment(v string) *GetAppResponseBody {
	s.RevisionComment = &v
	return s
}

func (s *GetAppResponseBody) SetRevisionTag(v string) *GetAppResponseBody {
	s.RevisionTag = &v
	return s
}

func (s *GetAppResponseBody) SetRevisions(v []*GetAppResponseBodyRevisions) *GetAppResponseBody {
	s.Revisions = v
	return s
}

func (s *GetAppResponseBody) SetScope(v string) *GetAppResponseBody {
	s.Scope = &v
	return s
}

func (s *GetAppResponseBody) SetSource(v string) *GetAppResponseBody {
	s.Source = &v
	return s
}

func (s *GetAppResponseBody) SetURL(v string) *GetAppResponseBody {
	s.URL = &v
	return s
}

func (s *GetAppResponseBody) SetWorkflowName(v string) *GetAppResponseBody {
	s.WorkflowName = &v
	return s
}

func (s *GetAppResponseBody) SetWorkspace(v string) *GetAppResponseBody {
	s.Workspace = &v
	return s
}

type GetAppResponseBodyConfigs struct {
	// example:
	//
	// {     "wgs.apply_bqsr.cpu": 4,     "wgs.apply_bqsr.disks": "local-disk 250 cloud_ssd",     "wgs.apply_bqsr.gatk_Launcher": "/usr/local/bin/gatk-4.1.4.1/gatk",     "wgs.apply_bqsr.java_opts": "\\"-XX:GCTimeLimit=50 -XX:GCHeapFreeLimit=10 -Xms3000m -Djava.io.tmpdir=/mnt\\"",     "wgs.apply_bqsr.memory": "8G"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// ./config/example.json
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s GetAppResponseBodyConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyConfigs) SetContent(v string) *GetAppResponseBodyConfigs {
	s.Content = &v
	return s
}

func (s *GetAppResponseBodyConfigs) SetPath(v string) *GetAppResponseBodyConfigs {
	s.Path = &v
	return s
}

type GetAppResponseBodyDependencies struct {
	// example:
	//
	// task echo {   String txt_out = "default_txt_out"   String dd_out = "default_dd_out"    command {     echo Hello > ${txt_out}     /bin/bash -c "dd if=/dev/zero of=${dd_out} bs=1M count=15"     echo "This is a test std out."   }    runtime {     cpu: "2"     memory: "4G"     autoReleaseJob: false     reserveOnFail: true     userData: "key1 value1"     disks: "local-disk 40 cloud_ssd, /home/mount/ 40 cloud_efficiency"     tag: "TestTag"   }    output {     File outTxtFile = "${txt_out}"     File outDdFile = "${dd_out}"   } }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// ./task/echo.wdl
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s GetAppResponseBodyDependencies) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyDependencies) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyDependencies) SetContent(v string) *GetAppResponseBodyDependencies {
	s.Content = &v
	return s
}

func (s *GetAppResponseBodyDependencies) SetPath(v string) *GetAppResponseBodyDependencies {
	s.Path = &v
	return s
}

type GetAppResponseBodyInputs struct {
	// example:
	//
	// help information
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// 1
	StepOrder *int64 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	// example:
	//
	// TestTaskName
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// TestVariableName
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// example:
	//
	// TestVariableType
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// example:
	//
	// TestVariableValue
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s GetAppResponseBodyInputs) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyInputs) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyInputs) SetHelp(v string) *GetAppResponseBodyInputs {
	s.Help = &v
	return s
}

func (s *GetAppResponseBodyInputs) SetRequired(v bool) *GetAppResponseBodyInputs {
	s.Required = &v
	return s
}

func (s *GetAppResponseBodyInputs) SetStepOrder(v int64) *GetAppResponseBodyInputs {
	s.StepOrder = &v
	return s
}

func (s *GetAppResponseBodyInputs) SetTaskName(v string) *GetAppResponseBodyInputs {
	s.TaskName = &v
	return s
}

func (s *GetAppResponseBodyInputs) SetVariableName(v string) *GetAppResponseBodyInputs {
	s.VariableName = &v
	return s
}

func (s *GetAppResponseBodyInputs) SetVariableType(v string) *GetAppResponseBodyInputs {
	s.VariableType = &v
	return s
}

func (s *GetAppResponseBodyInputs) SetVariableValue(v string) *GetAppResponseBodyInputs {
	s.VariableValue = &v
	return s
}

type GetAppResponseBodyOutputs struct {
	// example:
	//
	// Example help information
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// 1
	StepOrder *int64 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	// example:
	//
	// TestTaskName
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// TestVariableName
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// example:
	//
	// Int
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// example:
	//
	// ExampleValue
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s GetAppResponseBodyOutputs) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyOutputs) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyOutputs) SetHelp(v string) *GetAppResponseBodyOutputs {
	s.Help = &v
	return s
}

func (s *GetAppResponseBodyOutputs) SetRequired(v bool) *GetAppResponseBodyOutputs {
	s.Required = &v
	return s
}

func (s *GetAppResponseBodyOutputs) SetStepOrder(v int64) *GetAppResponseBodyOutputs {
	s.StepOrder = &v
	return s
}

func (s *GetAppResponseBodyOutputs) SetTaskName(v string) *GetAppResponseBodyOutputs {
	s.TaskName = &v
	return s
}

func (s *GetAppResponseBodyOutputs) SetVariableName(v string) *GetAppResponseBodyOutputs {
	s.VariableName = &v
	return s
}

func (s *GetAppResponseBodyOutputs) SetVariableType(v string) *GetAppResponseBodyOutputs {
	s.VariableType = &v
	return s
}

func (s *GetAppResponseBodyOutputs) SetVariableValue(v string) *GetAppResponseBodyOutputs {
	s.VariableValue = &v
	return s
}

type GetAppResponseBodyRevisions struct {
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// example:
	//
	// Example revision comment
	RevisionComment *string `json:"RevisionComment,omitempty" xml:"RevisionComment,omitempty"`
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
}

func (s GetAppResponseBodyRevisions) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyRevisions) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyRevisions) SetCreateTime(v string) *GetAppResponseBodyRevisions {
	s.CreateTime = &v
	return s
}

func (s *GetAppResponseBodyRevisions) SetRevision(v string) *GetAppResponseBodyRevisions {
	s.Revision = &v
	return s
}

func (s *GetAppResponseBodyRevisions) SetRevisionComment(v string) *GetAppResponseBodyRevisions {
	s.RevisionComment = &v
	return s
}

func (s *GetAppResponseBodyRevisions) SetRevisionTag(v string) *GetAppResponseBodyRevisions {
	s.RevisionTag = &v
	return s
}

type GetAppResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponse) GoString() string {
	return s.String()
}

func (s *GetAppResponse) SetHeaders(v map[string]*string) *GetAppResponse {
	s.Headers = v
	return s
}

func (s *GetAppResponse) SetStatusCode(v int32) *GetAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppResponse) SetBody(v *GetAppResponseBody) *GetAppResponse {
	s.Body = v
	return s
}

type GetEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRequest) GoString() string {
	return s.String()
}

func (s *GetEntityRequest) SetEntityType(v string) *GetEntityRequest {
	s.EntityType = &v
	return s
}

func (s *GetEntityRequest) SetWorkspace(v string) *GetEntityRequest {
	s.Workspace = &v
	return s
}

type GetEntityResponseBody struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityResponseBody) SetAttributes(v []*string) *GetEntityResponseBody {
	s.Attributes = v
	return s
}

func (s *GetEntityResponseBody) SetEntityType(v string) *GetEntityResponseBody {
	s.EntityType = &v
	return s
}

func (s *GetEntityResponseBody) SetHostId(v string) *GetEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *GetEntityResponseBody) SetRequestId(v string) *GetEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntityResponseBody) SetTotalCount(v int32) *GetEntityResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetEntityResponseBody) SetWorkspace(v string) *GetEntityResponseBody {
	s.Workspace = &v
	return s
}

type GetEntityResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEntityResponse) GoString() string {
	return s.String()
}

func (s *GetEntityResponse) SetHeaders(v map[string]*string) *GetEntityResponse {
	s.Headers = v
	return s
}

func (s *GetEntityResponse) SetStatusCode(v int32) *GetEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEntityResponse) SetBody(v *GetEntityResponseBody) *GetEntityResponse {
	s.Body = v
	return s
}

type GetGlobalAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// joint-calling
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// AppVersion, AppDescriptorFiles, DescriptorType, DAG, Document, Comment, Alias, NamespaceName, AppScope, AppName, AppDescription, RegionIds, Categories, Toolkit, AppVersions, AppDefaultVersion, AppFee, LastModified, Pinned, Contact, Links, AppType
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// cn-shenzhen
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sentieon
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s GetGlobalAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalAppRequest) GoString() string {
	return s.String()
}

func (s *GetGlobalAppRequest) SetAppName(v string) *GetGlobalAppRequest {
	s.AppName = &v
	return s
}

func (s *GetGlobalAppRequest) SetAppVersion(v string) *GetGlobalAppRequest {
	s.AppVersion = &v
	return s
}

func (s *GetGlobalAppRequest) SetAttributes(v []*string) *GetGlobalAppRequest {
	s.Attributes = v
	return s
}

func (s *GetGlobalAppRequest) SetLocation(v string) *GetGlobalAppRequest {
	s.Location = &v
	return s
}

func (s *GetGlobalAppRequest) SetNamespaceName(v string) *GetGlobalAppRequest {
	s.NamespaceName = &v
	return s
}

type GetGlobalAppShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// joint-calling
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// AppVersion, AppDescriptorFiles, DescriptorType, DAG, Document, Comment, Alias, NamespaceName, AppScope, AppName, AppDescription, RegionIds, Categories, Toolkit, AppVersions, AppDefaultVersion, AppFee, LastModified, Pinned, Contact, Links, AppType
	AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// cn-shenzhen
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sentieon
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s GetGlobalAppShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetGlobalAppShrinkRequest) SetAppName(v string) *GetGlobalAppShrinkRequest {
	s.AppName = &v
	return s
}

func (s *GetGlobalAppShrinkRequest) SetAppVersion(v string) *GetGlobalAppShrinkRequest {
	s.AppVersion = &v
	return s
}

func (s *GetGlobalAppShrinkRequest) SetAttributesShrink(v string) *GetGlobalAppShrinkRequest {
	s.AttributesShrink = &v
	return s
}

func (s *GetGlobalAppShrinkRequest) SetLocation(v string) *GetGlobalAppShrinkRequest {
	s.Location = &v
	return s
}

func (s *GetGlobalAppShrinkRequest) SetNamespaceName(v string) *GetGlobalAppShrinkRequest {
	s.NamespaceName = &v
	return s
}

type GetGlobalAppResponseBody struct {
	// example:
	//
	// 1.0
	AppDefaultVersion *string `json:"AppDefaultVersion,omitempty" xml:"AppDefaultVersion,omitempty"`
	// example:
	//
	// This is app description.
	AppDescription     *string                                       `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	AppDescriptorFiles []*GetGlobalAppResponseBodyAppDescriptorFiles `json:"AppDescriptorFiles,omitempty" xml:"AppDescriptorFiles,omitempty" type:"Repeated"`
	// example:
	//
	// WDL / CWL
	AppDescriptorType *string `json:"AppDescriptorType,omitempty" xml:"AppDescriptorType,omitempty"`
	// example:
	//
	// 实际使用时，计算该应用消耗的资源费用
	AppFee *string `json:"AppFee,omitempty" xml:"AppFee,omitempty"`
	// example:
	//
	// joint-calling
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// OpenAccess / Protected
	AppScope *string `json:"AppScope,omitempty" xml:"AppScope,omitempty"`
	// example:
	//
	// workflow
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 1.2a-3
	AppVersion  *string                                `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AppVersions []*GetGlobalAppResponseBodyAppVersions `json:"AppVersions,omitempty" xml:"AppVersions,omitempty" type:"Repeated"`
	Categories  []*string                              `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// example:
	//
	// 初始版本
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// support@sentieon.com
	Contact *string `json:"Contact,omitempty" xml:"Contact,omitempty"`
	// example:
	//
	// DAG
	DAG *string `json:"DAG,omitempty" xml:"DAG,omitempty"`
	// example:
	//
	// The app details.
	Document *string `json:"Document,omitempty" xml:"Document,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// yyyy/MM/dd HH:mm:ss
	LastModified *string   `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Links        []*string `json:"Links,omitempty" xml:"Links,omitempty" type:"Repeated"`
	Locations    []*string `json:"Locations,omitempty" xml:"Locations,omitempty" type:"Repeated"`
	// example:
	//
	// sentieon
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// example:
	//
	// true
	Pinned *bool `json:"Pinned,omitempty" xml:"Pinned,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Sentieon
	Toolkit *string `json:"Toolkit,omitempty" xml:"Toolkit,omitempty"`
}

func (s GetGlobalAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetGlobalAppResponseBody) SetAppDefaultVersion(v string) *GetGlobalAppResponseBody {
	s.AppDefaultVersion = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppDescription(v string) *GetGlobalAppResponseBody {
	s.AppDescription = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppDescriptorFiles(v []*GetGlobalAppResponseBodyAppDescriptorFiles) *GetGlobalAppResponseBody {
	s.AppDescriptorFiles = v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppDescriptorType(v string) *GetGlobalAppResponseBody {
	s.AppDescriptorType = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppFee(v string) *GetGlobalAppResponseBody {
	s.AppFee = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppName(v string) *GetGlobalAppResponseBody {
	s.AppName = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppScope(v string) *GetGlobalAppResponseBody {
	s.AppScope = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppType(v string) *GetGlobalAppResponseBody {
	s.AppType = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppVersion(v string) *GetGlobalAppResponseBody {
	s.AppVersion = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppVersions(v []*GetGlobalAppResponseBodyAppVersions) *GetGlobalAppResponseBody {
	s.AppVersions = v
	return s
}

func (s *GetGlobalAppResponseBody) SetCategories(v []*string) *GetGlobalAppResponseBody {
	s.Categories = v
	return s
}

func (s *GetGlobalAppResponseBody) SetComment(v string) *GetGlobalAppResponseBody {
	s.Comment = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetContact(v string) *GetGlobalAppResponseBody {
	s.Contact = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetDAG(v string) *GetGlobalAppResponseBody {
	s.DAG = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetDocument(v string) *GetGlobalAppResponseBody {
	s.Document = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetHostId(v string) *GetGlobalAppResponseBody {
	s.HostId = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetLastModified(v string) *GetGlobalAppResponseBody {
	s.LastModified = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetLinks(v []*string) *GetGlobalAppResponseBody {
	s.Links = v
	return s
}

func (s *GetGlobalAppResponseBody) SetLocations(v []*string) *GetGlobalAppResponseBody {
	s.Locations = v
	return s
}

func (s *GetGlobalAppResponseBody) SetNamespaceName(v string) *GetGlobalAppResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetPinned(v bool) *GetGlobalAppResponseBody {
	s.Pinned = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetRequestId(v string) *GetGlobalAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetToolkit(v string) *GetGlobalAppResponseBody {
	s.Toolkit = &v
	return s
}

type GetGlobalAppResponseBodyAppDescriptorFiles struct {
	Checksum *string `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// PRIMARY_DESCRIPTOR（主WDL） SECONDARY_DESCRIPTOR（依赖WDL）  CONFIG（推荐的输入参数）  CONTAINERFILE（镜像制作文件）  DATA（输入参数中引用的参考文件）  OTHER（其他）
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// ./main.wdl
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetGlobalAppResponseBodyAppDescriptorFiles) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalAppResponseBodyAppDescriptorFiles) GoString() string {
	return s.String()
}

func (s *GetGlobalAppResponseBodyAppDescriptorFiles) SetChecksum(v string) *GetGlobalAppResponseBodyAppDescriptorFiles {
	s.Checksum = &v
	return s
}

func (s *GetGlobalAppResponseBodyAppDescriptorFiles) SetContent(v string) *GetGlobalAppResponseBodyAppDescriptorFiles {
	s.Content = &v
	return s
}

func (s *GetGlobalAppResponseBodyAppDescriptorFiles) SetFileType(v string) *GetGlobalAppResponseBodyAppDescriptorFiles {
	s.FileType = &v
	return s
}

func (s *GetGlobalAppResponseBodyAppDescriptorFiles) SetPath(v string) *GetGlobalAppResponseBodyAppDescriptorFiles {
	s.Path = &v
	return s
}

func (s *GetGlobalAppResponseBodyAppDescriptorFiles) SetUrl(v string) *GetGlobalAppResponseBodyAppDescriptorFiles {
	s.Url = &v
	return s
}

type GetGlobalAppResponseBodyAppVersions struct {
	// example:
	//
	// 1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 初始版本
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// yyyy/MM/dd HH:mm:ss
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s GetGlobalAppResponseBodyAppVersions) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalAppResponseBodyAppVersions) GoString() string {
	return s.String()
}

func (s *GetGlobalAppResponseBodyAppVersions) SetAppVersion(v string) *GetGlobalAppResponseBodyAppVersions {
	s.AppVersion = &v
	return s
}

func (s *GetGlobalAppResponseBodyAppVersions) SetComment(v string) *GetGlobalAppResponseBodyAppVersions {
	s.Comment = &v
	return s
}

func (s *GetGlobalAppResponseBodyAppVersions) SetLastModified(v string) *GetGlobalAppResponseBodyAppVersions {
	s.LastModified = &v
	return s
}

type GetGlobalAppResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGlobalAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGlobalAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalAppResponse) GoString() string {
	return s.String()
}

func (s *GetGlobalAppResponse) SetHeaders(v map[string]*string) *GetGlobalAppResponse {
	s.Headers = v
	return s
}

func (s *GetGlobalAppResponse) SetStatusCode(v int32) *GetGlobalAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGlobalAppResponse) SetBody(v *GetGlobalAppResponseBody) *GetGlobalAppResponse {
	s.Body = v
	return s
}

type GetPublicDatasetRequest struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 西蒙斯基因组多样性计划
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
}

func (s GetPublicDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetRequest) SetAttributes(v []*string) *GetPublicDatasetRequest {
	s.Attributes = v
	return s
}

func (s *GetPublicDatasetRequest) SetDatasetName(v string) *GetPublicDatasetRequest {
	s.DatasetName = &v
	return s
}

type GetPublicDatasetShrinkRequest struct {
	AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 西蒙斯基因组多样性计划
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
}

func (s GetPublicDatasetShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetShrinkRequest) SetAttributesShrink(v string) *GetPublicDatasetShrinkRequest {
	s.AttributesShrink = &v
	return s
}

func (s *GetPublicDatasetShrinkRequest) SetDatasetName(v string) *GetPublicDatasetShrinkRequest {
	s.DatasetName = &v
	return s
}

type GetPublicDatasetResponseBody struct {
	// example:
	//
	// This is description about dataset.
	About *string `json:"About,omitempty" xml:"About,omitempty"`
	// example:
	//
	// This are dataset access requirements.
	AccessRequirements *string `json:"AccessRequirements,omitempty" xml:"AccessRequirements,omitempty"`
	// example:
	//
	// This is dataset copyright.
	Copyright *string `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	// example:
	//
	// This is dataset description.
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// example:
	//
	// 西蒙斯基因组多样性计划
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// yyyy/MM/dd HH:mm:ss
	LastModified *string   `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Locations    []*string `json:"Locations,omitempty" xml:"Locations,omitempty" type:"Repeated"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 源站更新后及时同步
	UpdateFrequency *string `json:"UpdateFrequency,omitempty" xml:"UpdateFrequency,omitempty"`
}

func (s GetPublicDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetResponseBody) SetAbout(v string) *GetPublicDatasetResponseBody {
	s.About = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetAccessRequirements(v string) *GetPublicDatasetResponseBody {
	s.AccessRequirements = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetCopyright(v string) *GetPublicDatasetResponseBody {
	s.Copyright = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetDatasetDescription(v string) *GetPublicDatasetResponseBody {
	s.DatasetDescription = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetDatasetName(v string) *GetPublicDatasetResponseBody {
	s.DatasetName = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetHostId(v string) *GetPublicDatasetResponseBody {
	s.HostId = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetLastModified(v string) *GetPublicDatasetResponseBody {
	s.LastModified = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetLocations(v []*string) *GetPublicDatasetResponseBody {
	s.Locations = v
	return s
}

func (s *GetPublicDatasetResponseBody) SetRequestId(v string) *GetPublicDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetTags(v []*string) *GetPublicDatasetResponseBody {
	s.Tags = v
	return s
}

func (s *GetPublicDatasetResponseBody) SetUpdateFrequency(v string) *GetPublicDatasetResponseBody {
	s.UpdateFrequency = &v
	return s
}

type GetPublicDatasetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublicDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublicDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetResponse) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetResponse) SetHeaders(v map[string]*string) *GetPublicDatasetResponse {
	s.Headers = v
	return s
}

func (s *GetPublicDatasetResponse) SetStatusCode(v int32) *GetPublicDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublicDatasetResponse) SetBody(v *GetPublicDatasetResponseBody) *GetPublicDatasetResponse {
	s.Body = v
	return s
}

type GetPublicDatasetEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 西蒙斯基因组多样性计划
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s GetPublicDatasetEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetEntityRequest) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetEntityRequest) SetDatasetName(v string) *GetPublicDatasetEntityRequest {
	s.DatasetName = &v
	return s
}

func (s *GetPublicDatasetEntityRequest) SetEntityType(v string) *GetPublicDatasetEntityRequest {
	s.EntityType = &v
	return s
}

func (s *GetPublicDatasetEntityRequest) SetLocation(v string) *GetPublicDatasetEntityRequest {
	s.Location = &v
	return s
}

type GetPublicDatasetEntityResponseBody struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// 西蒙斯基因组多样性计划
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetPublicDatasetEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetEntityResponseBody) SetAttributes(v []*string) *GetPublicDatasetEntityResponseBody {
	s.Attributes = v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetDatasetName(v string) *GetPublicDatasetEntityResponseBody {
	s.DatasetName = &v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetEntityType(v string) *GetPublicDatasetEntityResponseBody {
	s.EntityType = &v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetHostId(v string) *GetPublicDatasetEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetRequestId(v string) *GetPublicDatasetEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetTotalCount(v int32) *GetPublicDatasetEntityResponseBody {
	s.TotalCount = &v
	return s
}

type GetPublicDatasetEntityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublicDatasetEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublicDatasetEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetEntityResponse) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetEntityResponse) SetHeaders(v map[string]*string) *GetPublicDatasetEntityResponse {
	s.Headers = v
	return s
}

func (s *GetPublicDatasetEntityResponse) SetStatusCode(v int32) *GetPublicDatasetEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublicDatasetEntityResponse) SetBody(v *GetPublicDatasetEntityResponseBody) *GetPublicDatasetEntityResponse {
	s.Body = v
	return s
}

type GetRunRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// run-12kDbbqp5O9dfsH0sEdQeRclGQb
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetRunRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRunRequest) GoString() string {
	return s.String()
}

func (s *GetRunRequest) SetRunId(v string) *GetRunRequest {
	s.RunId = &v
	return s
}

func (s *GetRunRequest) SetWorkspace(v string) *GetRunRequest {
	s.Workspace = &v
	return s
}

type GetRunResponseBody struct {
	// example:
	//
	// Sentieon
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	AppRevision        *string   `json:"AppRevision,omitempty" xml:"AppRevision,omitempty"`
	BillingInstanceIds []*string `json:"BillingInstanceIds,omitempty" xml:"BillingInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Calls *string `json:"Calls,omitempty" xml:"Calls,omitempty"`
	// example:
	//
	// 2021-06-10T13:30:20.414557061Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// {}
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	// example:
	//
	// test run
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2021-06-10T13:36:20.405878292Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// s1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// oss://my-bucket/exec/
	ExecuteDirectory *string                           `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	ExecuteOptions   *GetRunResponseBodyExecuteOptions `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty" type:"Struct"`
	// example:
	//
	// {}
	Failures *string `json:"Failures,omitempty" xml:"Failures,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// {     "workflow1.task1.param1": "helloooo.txt",     "workflow1.task1.param2": "abc",     "workflow1.task2.param1": "new_helloooo.txt",     "workflow1.task2.param2": "def" }
	Inputs *string            `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// oss://my-bucket/output/
	OutputFolder *string `json:"OutputFolder,omitempty" xml:"OutputFolder,omitempty"`
	// example:
	//
	// {         "workflow1.bam": "oss://my-bucket/test.bam",         "workflow1.vcf": "oss://my-bucket/test.vcf"     }
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// run-13BeYuxc0fxO24uA75UppTiMRoQ
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// example:
	//
	// test-run
	RunName *string `json:"RunName,omitempty" xml:"RunName,omitempty"`
	// example:
	//
	// EasyGene
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 2021-06-10T13:30:24.414557061Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// sub-13BeXWnGvrOl3MC124tazcJVKaV
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// example:
	//
	// {}
	Timing *string `json:"Timing,omitempty" xml:"Timing,omitempty"`
	// example:
	//
	// 13579222
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunResponseBody) SetAppName(v string) *GetRunResponseBody {
	s.AppName = &v
	return s
}

func (s *GetRunResponseBody) SetAppRevision(v string) *GetRunResponseBody {
	s.AppRevision = &v
	return s
}

func (s *GetRunResponseBody) SetBillingInstanceIds(v []*string) *GetRunResponseBody {
	s.BillingInstanceIds = v
	return s
}

func (s *GetRunResponseBody) SetCalls(v string) *GetRunResponseBody {
	s.Calls = &v
	return s
}

func (s *GetRunResponseBody) SetCreateTime(v string) *GetRunResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRunResponseBody) SetDefaultRuntime(v string) *GetRunResponseBody {
	s.DefaultRuntime = &v
	return s
}

func (s *GetRunResponseBody) SetDescription(v string) *GetRunResponseBody {
	s.Description = &v
	return s
}

func (s *GetRunResponseBody) SetEndTime(v string) *GetRunResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetRunResponseBody) SetEntityName(v string) *GetRunResponseBody {
	s.EntityName = &v
	return s
}

func (s *GetRunResponseBody) SetEntityType(v string) *GetRunResponseBody {
	s.EntityType = &v
	return s
}

func (s *GetRunResponseBody) SetExecuteDirectory(v string) *GetRunResponseBody {
	s.ExecuteDirectory = &v
	return s
}

func (s *GetRunResponseBody) SetExecuteOptions(v *GetRunResponseBodyExecuteOptions) *GetRunResponseBody {
	s.ExecuteOptions = v
	return s
}

func (s *GetRunResponseBody) SetFailures(v string) *GetRunResponseBody {
	s.Failures = &v
	return s
}

func (s *GetRunResponseBody) SetHostId(v string) *GetRunResponseBody {
	s.HostId = &v
	return s
}

func (s *GetRunResponseBody) SetInputs(v string) *GetRunResponseBody {
	s.Inputs = &v
	return s
}

func (s *GetRunResponseBody) SetLabels(v map[string]*string) *GetRunResponseBody {
	s.Labels = v
	return s
}

func (s *GetRunResponseBody) SetOutputFolder(v string) *GetRunResponseBody {
	s.OutputFolder = &v
	return s
}

func (s *GetRunResponseBody) SetOutputs(v string) *GetRunResponseBody {
	s.Outputs = &v
	return s
}

func (s *GetRunResponseBody) SetRequestId(v string) *GetRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRunResponseBody) SetRunId(v string) *GetRunResponseBody {
	s.RunId = &v
	return s
}

func (s *GetRunResponseBody) SetRunName(v string) *GetRunResponseBody {
	s.RunName = &v
	return s
}

func (s *GetRunResponseBody) SetSource(v string) *GetRunResponseBody {
	s.Source = &v
	return s
}

func (s *GetRunResponseBody) SetStartTime(v string) *GetRunResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetRunResponseBody) SetStatus(v string) *GetRunResponseBody {
	s.Status = &v
	return s
}

func (s *GetRunResponseBody) SetSubmissionId(v string) *GetRunResponseBody {
	s.SubmissionId = &v
	return s
}

func (s *GetRunResponseBody) SetTiming(v string) *GetRunResponseBody {
	s.Timing = &v
	return s
}

func (s *GetRunResponseBody) SetUser(v string) *GetRunResponseBody {
	s.User = &v
	return s
}

func (s *GetRunResponseBody) SetWorkspace(v string) *GetRunResponseBody {
	s.Workspace = &v
	return s
}

type GetRunResponseBodyExecuteOptions struct {
	// example:
	//
	// false
	CallCaching *bool   `json:"CallCaching,omitempty" xml:"CallCaching,omitempty"`
	DataLoader  *string `json:"DataLoader,omitempty" xml:"DataLoader,omitempty"`
	// example:
	//
	// false
	DeleteIntermediateResults *bool `json:"DeleteIntermediateResults,omitempty" xml:"DeleteIntermediateResults,omitempty"`
	// example:
	//
	// NoNewCalls
	FailureMode *string `json:"FailureMode,omitempty" xml:"FailureMode,omitempty"`
	// example:
	//
	// false
	UseRelativeOutputPaths *bool `json:"UseRelativeOutputPaths,omitempty" xml:"UseRelativeOutputPaths,omitempty"`
}

func (s GetRunResponseBodyExecuteOptions) String() string {
	return tea.Prettify(s)
}

func (s GetRunResponseBodyExecuteOptions) GoString() string {
	return s.String()
}

func (s *GetRunResponseBodyExecuteOptions) SetCallCaching(v bool) *GetRunResponseBodyExecuteOptions {
	s.CallCaching = &v
	return s
}

func (s *GetRunResponseBodyExecuteOptions) SetDataLoader(v string) *GetRunResponseBodyExecuteOptions {
	s.DataLoader = &v
	return s
}

func (s *GetRunResponseBodyExecuteOptions) SetDeleteIntermediateResults(v bool) *GetRunResponseBodyExecuteOptions {
	s.DeleteIntermediateResults = &v
	return s
}

func (s *GetRunResponseBodyExecuteOptions) SetFailureMode(v string) *GetRunResponseBodyExecuteOptions {
	s.FailureMode = &v
	return s
}

func (s *GetRunResponseBodyExecuteOptions) SetUseRelativeOutputPaths(v bool) *GetRunResponseBodyExecuteOptions {
	s.UseRelativeOutputPaths = &v
	return s
}

type GetRunResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRunResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRunResponse) GoString() string {
	return s.String()
}

func (s *GetRunResponse) SetHeaders(v map[string]*string) *GetRunResponse {
	s.Headers = v
	return s
}

func (s *GetRunResponse) SetStatusCode(v int32) *GetRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRunResponse) SetBody(v *GetRunResponseBody) *GetRunResponse {
	s.Body = v
	return s
}

type GetSubmissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sub-13BeXWnGvrOl3MC124tazcJVKaV
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetSubmissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubmissionRequest) GoString() string {
	return s.String()
}

func (s *GetSubmissionRequest) SetSubmissionId(v string) *GetSubmissionRequest {
	s.SubmissionId = &v
	return s
}

func (s *GetSubmissionRequest) SetWorkspace(v string) *GetSubmissionRequest {
	s.Workspace = &v
	return s
}

type GetSubmissionResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Submission *GetSubmissionResponseBodySubmission `json:"Submission,omitempty" xml:"Submission,omitempty" type:"Struct"`
}

func (s GetSubmissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubmissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubmissionResponseBody) SetHostId(v string) *GetSubmissionResponseBody {
	s.HostId = &v
	return s
}

func (s *GetSubmissionResponseBody) SetRequestId(v string) *GetSubmissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubmissionResponseBody) SetSubmission(v *GetSubmissionResponseBodySubmission) *GetSubmissionResponseBody {
	s.Submission = v
	return s
}

type GetSubmissionResponseBodySubmission struct {
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// TestEntityType
	EntityType *string                                      `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	RunStats   *GetSubmissionResponseBodySubmissionRunStats `json:"RunStats,omitempty" xml:"RunStats,omitempty" type:"Struct"`
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// sub-13BeXWnGvrOl3MC124tazcJVKaV
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetSubmissionResponseBodySubmission) String() string {
	return tea.Prettify(s)
}

func (s GetSubmissionResponseBodySubmission) GoString() string {
	return s.String()
}

func (s *GetSubmissionResponseBodySubmission) SetCreateTime(v string) *GetSubmissionResponseBodySubmission {
	s.CreateTime = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetEndTime(v string) *GetSubmissionResponseBodySubmission {
	s.EndTime = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetEntityType(v string) *GetSubmissionResponseBodySubmission {
	s.EntityType = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetRunStats(v *GetSubmissionResponseBodySubmissionRunStats) *GetSubmissionResponseBodySubmission {
	s.RunStats = v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetStartTime(v string) *GetSubmissionResponseBodySubmission {
	s.StartTime = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetStatus(v string) *GetSubmissionResponseBodySubmission {
	s.Status = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetSubmissionId(v string) *GetSubmissionResponseBodySubmission {
	s.SubmissionId = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetWorkspace(v string) *GetSubmissionResponseBodySubmission {
	s.Workspace = &v
	return s
}

type GetSubmissionResponseBodySubmissionRunStats struct {
	// example:
	//
	// 5
	Aborted *int64 `json:"Aborted,omitempty" xml:"Aborted,omitempty"`
	// example:
	//
	// 5
	Aborting *int64 `json:"Aborting,omitempty" xml:"Aborting,omitempty"`
	// example:
	//
	// 5
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// example:
	//
	// 5
	Pending *int64 `json:"Pending,omitempty" xml:"Pending,omitempty"`
	// example:
	//
	// 5
	Running *int64 `json:"Running,omitempty" xml:"Running,omitempty"`
	// example:
	//
	// 5
	Submitted *int64 `json:"Submitted,omitempty" xml:"Submitted,omitempty"`
	// example:
	//
	// 5
	Succeeded *int64 `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s GetSubmissionResponseBodySubmissionRunStats) String() string {
	return tea.Prettify(s)
}

func (s GetSubmissionResponseBodySubmissionRunStats) GoString() string {
	return s.String()
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetAborted(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Aborted = &v
	return s
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetAborting(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Aborting = &v
	return s
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetFailed(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Failed = &v
	return s
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetPending(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Pending = &v
	return s
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetRunning(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Running = &v
	return s
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetSubmitted(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Submitted = &v
	return s
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetSucceeded(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Succeeded = &v
	return s
}

type GetSubmissionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubmissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubmissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubmissionResponse) GoString() string {
	return s.String()
}

func (s *GetSubmissionResponse) SetHeaders(v map[string]*string) *GetSubmissionResponse {
	s.Headers = v
	return s
}

func (s *GetSubmissionResponse) SetStatusCode(v int32) *GetSubmissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubmissionResponse) SetBody(v *GetSubmissionResponseBody) *GetSubmissionResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetTemplateName(v string) *GetTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateRequest) SetWorkspace(v string) *GetTemplateRequest {
	s.Workspace = &v
	return s
}

type GetTemplateResponseBody struct {
	// example:
	//
	// MyApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	AppRevision *string `json:"AppRevision,omitempty" xml:"AppRevision,omitempty"`
	// example:
	//
	// 2021-06-10T13:30:20.414557061Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId           *string                                    `json:"HostId,omitempty" xml:"HostId,omitempty"`
	InputsExpression []*GetTemplateResponseBodyInputsExpression `json:"InputsExpression,omitempty" xml:"InputsExpression,omitempty" type:"Repeated"`
	Labels           map[string]*string                         `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// 2021-06-10T13:30:20.405878292Z
	LastModifiedTime  *string                                     `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	OutputsExpression []*GetTemplateResponseBodyOutputsExpression `json:"OutputsExpression,omitempty" xml:"OutputsExpression,omitempty" type:"Repeated"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// produceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// example:
	//
	// test-entity
	RootEntity *string `json:"RootEntity,omitempty" xml:"RootEntity,omitempty"`
	// example:
	//
	// EasyGene
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// test-template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetAppName(v string) *GetTemplateResponseBody {
	s.AppName = &v
	return s
}

func (s *GetTemplateResponseBody) SetAppRevision(v string) *GetTemplateResponseBody {
	s.AppRevision = &v
	return s
}

func (s *GetTemplateResponseBody) SetCreateTime(v string) *GetTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateResponseBody) SetDescription(v string) *GetTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBody) SetHostId(v string) *GetTemplateResponseBody {
	s.HostId = &v
	return s
}

func (s *GetTemplateResponseBody) SetInputsExpression(v []*GetTemplateResponseBodyInputsExpression) *GetTemplateResponseBody {
	s.InputsExpression = v
	return s
}

func (s *GetTemplateResponseBody) SetLabels(v map[string]*string) *GetTemplateResponseBody {
	s.Labels = v
	return s
}

func (s *GetTemplateResponseBody) SetLastModifiedTime(v string) *GetTemplateResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetTemplateResponseBody) SetOutputsExpression(v []*GetTemplateResponseBodyOutputsExpression) *GetTemplateResponseBody {
	s.OutputsExpression = v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetRevisionTag(v string) *GetTemplateResponseBody {
	s.RevisionTag = &v
	return s
}

func (s *GetTemplateResponseBody) SetRootEntity(v string) *GetTemplateResponseBody {
	s.RootEntity = &v
	return s
}

func (s *GetTemplateResponseBody) SetSource(v string) *GetTemplateResponseBody {
	s.Source = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateName(v string) *GetTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateResponseBody) SetWorkspace(v string) *GetTemplateResponseBody {
	s.Workspace = &v
	return s
}

type GetTemplateResponseBodyInputsExpression struct {
	// example:
	//
	// test help
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// 0
	StepOrder *int64 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	// example:
	//
	// Task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// v1
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// example:
	//
	// File
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// example:
	//
	// oss://bucket/file1.fastq
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s GetTemplateResponseBodyInputsExpression) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyInputsExpression) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyInputsExpression) SetHelp(v string) *GetTemplateResponseBodyInputsExpression {
	s.Help = &v
	return s
}

func (s *GetTemplateResponseBodyInputsExpression) SetRequired(v bool) *GetTemplateResponseBodyInputsExpression {
	s.Required = &v
	return s
}

func (s *GetTemplateResponseBodyInputsExpression) SetStepOrder(v int64) *GetTemplateResponseBodyInputsExpression {
	s.StepOrder = &v
	return s
}

func (s *GetTemplateResponseBodyInputsExpression) SetTaskName(v string) *GetTemplateResponseBodyInputsExpression {
	s.TaskName = &v
	return s
}

func (s *GetTemplateResponseBodyInputsExpression) SetVariableName(v string) *GetTemplateResponseBodyInputsExpression {
	s.VariableName = &v
	return s
}

func (s *GetTemplateResponseBodyInputsExpression) SetVariableType(v string) *GetTemplateResponseBodyInputsExpression {
	s.VariableType = &v
	return s
}

func (s *GetTemplateResponseBodyInputsExpression) SetVariableValue(v string) *GetTemplateResponseBodyInputsExpression {
	s.VariableValue = &v
	return s
}

type GetTemplateResponseBodyOutputsExpression struct {
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// 2
	StepOrder *int64 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	// example:
	//
	// jobname
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// fastq1
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// example:
	//
	// string
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// example:
	//
	// abc_1.fastq.gz
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s GetTemplateResponseBodyOutputsExpression) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyOutputsExpression) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyOutputsExpression) SetHelp(v string) *GetTemplateResponseBodyOutputsExpression {
	s.Help = &v
	return s
}

func (s *GetTemplateResponseBodyOutputsExpression) SetRequired(v bool) *GetTemplateResponseBodyOutputsExpression {
	s.Required = &v
	return s
}

func (s *GetTemplateResponseBodyOutputsExpression) SetStepOrder(v int64) *GetTemplateResponseBodyOutputsExpression {
	s.StepOrder = &v
	return s
}

func (s *GetTemplateResponseBodyOutputsExpression) SetTaskName(v string) *GetTemplateResponseBodyOutputsExpression {
	s.TaskName = &v
	return s
}

func (s *GetTemplateResponseBodyOutputsExpression) SetVariableName(v string) *GetTemplateResponseBodyOutputsExpression {
	s.VariableName = &v
	return s
}

func (s *GetTemplateResponseBodyOutputsExpression) SetVariableType(v string) *GetTemplateResponseBodyOutputsExpression {
	s.VariableType = &v
	return s
}

func (s *GetTemplateResponseBodyOutputsExpression) SetVariableValue(v string) *GetTemplateResponseBodyOutputsExpression {
	s.VariableValue = &v
	return s
}

type GetTemplateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetStatusCode(v int32) *GetTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type GetWorkspaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// MyTestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRequest) SetWorkspace(v string) *GetWorkspaceRequest {
	s.Workspace = &v
	return s
}

type GetWorkspaceResponseBody struct {
	// example:
	//
	// my-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// This is a test workspace
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 30
	JobLifecycle *int32             `json:"JobLifecycle,omitempty" xml:"JobLifecycle,omitempty"`
	Labels       map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// example:
	//
	// cn-beijing
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// TestRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// oss://my-bucket/my-path/
	Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) SetBucketName(v string) *GetWorkspaceResponseBody {
	s.BucketName = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetCreateTime(v string) *GetWorkspaceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetDescription(v string) *GetWorkspaceResponseBody {
	s.Description = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetHostId(v string) *GetWorkspaceResponseBody {
	s.HostId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetJobLifecycle(v int32) *GetWorkspaceResponseBody {
	s.JobLifecycle = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetLabels(v map[string]*string) *GetWorkspaceResponseBody {
	s.Labels = v
	return s
}

func (s *GetWorkspaceResponseBody) SetLastModifiedTime(v string) *GetWorkspaceResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetLocation(v string) *GetWorkspaceResponseBody {
	s.Location = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRole(v string) *GetWorkspaceResponseBody {
	s.Role = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetStatus(v string) *GetWorkspaceResponseBody {
	s.Status = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetStorage(v string) *GetWorkspaceResponseBody {
	s.Storage = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspace(v string) *GetWorkspaceResponseBody {
	s.Workspace = &v
	return s
}

type GetWorkspaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponse) SetHeaders(v map[string]*string) *GetWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceResponse) SetStatusCode(v int32) *GetWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceResponse) SetBody(v *GetWorkspaceResponseBody) *GetWorkspaceResponse {
	s.Body = v
	return s
}

type ImportAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// fastp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// easygene:opengene/fastp
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ImportAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportAppRequest) GoString() string {
	return s.String()
}

func (s *ImportAppRequest) SetAppName(v string) *ImportAppRequest {
	s.AppName = &v
	return s
}

func (s *ImportAppRequest) SetSource(v string) *ImportAppRequest {
	s.Source = &v
	return s
}

func (s *ImportAppRequest) SetWorkspace(v string) *ImportAppRequest {
	s.Workspace = &v
	return s
}

type ImportAppResponseBody struct {
	// example:
	//
	// InstalledApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ImportAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportAppResponseBody) GoString() string {
	return s.String()
}

func (s *ImportAppResponseBody) SetAppName(v string) *ImportAppResponseBody {
	s.AppName = &v
	return s
}

func (s *ImportAppResponseBody) SetHostId(v string) *ImportAppResponseBody {
	s.HostId = &v
	return s
}

func (s *ImportAppResponseBody) SetRegionId(v string) *ImportAppResponseBody {
	s.RegionId = &v
	return s
}

func (s *ImportAppResponseBody) SetRequestId(v string) *ImportAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportAppResponseBody) SetWorkspace(v string) *ImportAppResponseBody {
	s.Workspace = &v
	return s
}

type ImportAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportAppResponse) GoString() string {
	return s.String()
}

func (s *ImportAppResponse) SetHeaders(v map[string]*string) *ImportAppResponse {
	s.Headers = v
	return s
}

func (s *ImportAppResponse) SetStatusCode(v int32) *ImportAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportAppResponse) SetBody(v *ImportAppResponseBody) *ImportAppResponse {
	s.Body = v
	return s
}

type InstallGlobalAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// joint-calling
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// joint-calling
	InstalledAppName *string `json:"InstalledAppName,omitempty" xml:"InstalledAppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sentieon
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EasyGene
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s InstallGlobalAppRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallGlobalAppRequest) GoString() string {
	return s.String()
}

func (s *InstallGlobalAppRequest) SetAppName(v string) *InstallGlobalAppRequest {
	s.AppName = &v
	return s
}

func (s *InstallGlobalAppRequest) SetInstalledAppName(v string) *InstallGlobalAppRequest {
	s.InstalledAppName = &v
	return s
}

func (s *InstallGlobalAppRequest) SetNamespaceName(v string) *InstallGlobalAppRequest {
	s.NamespaceName = &v
	return s
}

func (s *InstallGlobalAppRequest) SetSource(v string) *InstallGlobalAppRequest {
	s.Source = &v
	return s
}

func (s *InstallGlobalAppRequest) SetWorkspace(v string) *InstallGlobalAppRequest {
	s.Workspace = &v
	return s
}

type InstallGlobalAppResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// InstalledApp
	InstalledAppName *string `json:"InstalledAppName,omitempty" xml:"InstalledAppName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s InstallGlobalAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallGlobalAppResponseBody) GoString() string {
	return s.String()
}

func (s *InstallGlobalAppResponseBody) SetHostId(v string) *InstallGlobalAppResponseBody {
	s.HostId = &v
	return s
}

func (s *InstallGlobalAppResponseBody) SetInstalledAppName(v string) *InstallGlobalAppResponseBody {
	s.InstalledAppName = &v
	return s
}

func (s *InstallGlobalAppResponseBody) SetRegionId(v string) *InstallGlobalAppResponseBody {
	s.RegionId = &v
	return s
}

func (s *InstallGlobalAppResponseBody) SetRequestId(v string) *InstallGlobalAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallGlobalAppResponseBody) SetWorkspace(v string) *InstallGlobalAppResponseBody {
	s.Workspace = &v
	return s
}

type InstallGlobalAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallGlobalAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallGlobalAppResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallGlobalAppResponse) GoString() string {
	return s.String()
}

func (s *InstallGlobalAppResponse) SetHeaders(v map[string]*string) *InstallGlobalAppResponse {
	s.Headers = v
	return s
}

func (s *InstallGlobalAppResponse) SetStatusCode(v int32) *InstallGlobalAppResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallGlobalAppResponse) SetBody(v *InstallGlobalAppResponseBody) *InstallGlobalAppResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	// example:
	//
	// App
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// true
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// example:
	//
	// env=test
	LabelSelector *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	// example:
	//
	// WDL
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Next Token
	//
	// example:
	//
	// TestToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// createtime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// Public
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// TestName
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetAppType(v string) *ListAppsRequest {
	s.AppType = &v
	return s
}

func (s *ListAppsRequest) SetIsReversed(v bool) *ListAppsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListAppsRequest) SetLabelSelector(v string) *ListAppsRequest {
	s.LabelSelector = &v
	return s
}

func (s *ListAppsRequest) SetLanguage(v string) *ListAppsRequest {
	s.Language = &v
	return s
}

func (s *ListAppsRequest) SetMaxResults(v int32) *ListAppsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppsRequest) SetNextToken(v string) *ListAppsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppsRequest) SetOrderBy(v string) *ListAppsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListAppsRequest) SetScope(v string) *ListAppsRequest {
	s.Scope = &v
	return s
}

func (s *ListAppsRequest) SetSearch(v string) *ListAppsRequest {
	s.Search = &v
	return s
}

func (s *ListAppsRequest) SetWorkspace(v string) *ListAppsRequest {
	s.Workspace = &v
	return s
}

type ListAppsResponseBody struct {
	Apps []*ListAppsResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Next Token
	//
	// example:
	//
	// TestToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) SetApps(v []*ListAppsResponseBodyApps) *ListAppsResponseBody {
	s.Apps = v
	return s
}

func (s *ListAppsResponseBody) SetHostId(v string) *ListAppsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListAppsResponseBody) SetMaxResults(v int32) *ListAppsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppsResponseBody) SetNextToken(v string) *ListAppsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsResponseBody) SetTotalCount(v int32) *ListAppsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppsResponseBodyApps struct {
	// example:
	//
	// v1
	AppDefaultVersion *string `json:"AppDefaultVersion,omitempty" xml:"AppDefaultVersion,omitempty"`
	// example:
	//
	// TestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// App
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 2021-06-10T13:30:20.414557061Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// This is a test app
	Description *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels      map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// WDL
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// Public
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// EasyGene
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ListAppsResponseBodyApps) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyApps) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyApps) SetAppDefaultVersion(v string) *ListAppsResponseBodyApps {
	s.AppDefaultVersion = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetAppName(v string) *ListAppsResponseBodyApps {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetAppType(v string) *ListAppsResponseBodyApps {
	s.AppType = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetCreateTime(v string) *ListAppsResponseBodyApps {
	s.CreateTime = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetDescription(v string) *ListAppsResponseBodyApps {
	s.Description = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetLabels(v map[string]*string) *ListAppsResponseBodyApps {
	s.Labels = v
	return s
}

func (s *ListAppsResponseBodyApps) SetLanguage(v string) *ListAppsResponseBodyApps {
	s.Language = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetScope(v string) *ListAppsResponseBodyApps {
	s.Scope = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetSource(v string) *ListAppsResponseBodyApps {
	s.Source = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetWorkspace(v string) *ListAppsResponseBodyApps {
	s.Workspace = &v
	return s
}

type ListAppsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetStatusCode(v int32) *ListAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppsResponse) SetBody(v *ListAppsResponseBody) *ListAppsResponse {
	s.Body = v
	return s
}

type ListAuthorizedSoftwareRequest struct {
	// example:
	//
	// true
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0f2cc944-67db-4fb1-bb10-f136cda74bda
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// LastModified
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// Sentieon
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
}

func (s ListAuthorizedSoftwareRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizedSoftwareRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedSoftwareRequest) SetIsReversed(v bool) *ListAuthorizedSoftwareRequest {
	s.IsReversed = &v
	return s
}

func (s *ListAuthorizedSoftwareRequest) SetLocation(v string) *ListAuthorizedSoftwareRequest {
	s.Location = &v
	return s
}

func (s *ListAuthorizedSoftwareRequest) SetMaxResults(v int32) *ListAuthorizedSoftwareRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizedSoftwareRequest) SetNextToken(v string) *ListAuthorizedSoftwareRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizedSoftwareRequest) SetOrderBy(v string) *ListAuthorizedSoftwareRequest {
	s.OrderBy = &v
	return s
}

func (s *ListAuthorizedSoftwareRequest) SetSearch(v string) *ListAuthorizedSoftwareRequest {
	s.Search = &v
	return s
}

type ListAuthorizedSoftwareResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ba6d88e2-a6f2-4d2c-a2ae-68b76b078445
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Softwares []*ListAuthorizedSoftwareResponseBodySoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthorizedSoftwareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizedSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedSoftwareResponseBody) SetHostId(v string) *ListAuthorizedSoftwareResponseBody {
	s.HostId = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetMaxResults(v int32) *ListAuthorizedSoftwareResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetNextToken(v string) *ListAuthorizedSoftwareResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetRequestId(v string) *ListAuthorizedSoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetSoftwares(v []*ListAuthorizedSoftwareResponseBodySoftwares) *ListAuthorizedSoftwareResponseBody {
	s.Softwares = v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetTotalCount(v int32) *ListAuthorizedSoftwareResponseBody {
	s.TotalCount = &v
	return s
}

type ListAuthorizedSoftwareResponseBodySoftwares struct {
	// example:
	//
	// https://help.aliyun.com/document_detail/257750.html
	HelpLink *string `json:"HelpLink,omitempty" xml:"HelpLink,omitempty"`
	// example:
	//
	// yyyy/MM/dd HH:mm:ss
	LastModified *string   `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Locations    []*string `json:"Locations,omitempty" xml:"Locations,omitempty" type:"Repeated"`
	Promotion    *string   `json:"Promotion,omitempty" xml:"Promotion,omitempty"`
	// example:
	//
	// 201911
	SoftwareDefaultVersion *string `json:"SoftwareDefaultVersion,omitempty" xml:"SoftwareDefaultVersion,omitempty"`
	// example:
	//
	// NGS数据二级分析软件，提供基于CPU的加速方案，可替代BWA，GATK，HaplotypeCaller，Mutect和Mutect2等分析流程。包括DNASeq，DNAScope，TNScope等工具
	SoftwareDescription *string `json:"SoftwareDescription,omitempty" xml:"SoftwareDescription,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i2/O1CN01mmZovA1UJfeKQAYvF_!!6000000002497-2-tps-256-256.png
	SoftwareIcon *string `json:"SoftwareIcon,omitempty" xml:"SoftwareIcon,omitempty"`
	// example:
	//
	// 0.15
	SoftwareLicenseFee *float32 `json:"SoftwareLicenseFee,omitempty" xml:"SoftwareLicenseFee,omitempty"`
	// example:
	//
	// Sentieon加速软件
	SoftwareLongName *string `json:"SoftwareLongName,omitempty" xml:"SoftwareLongName,omitempty"`
	// example:
	//
	// sentieon
	SoftwareName     *string   `json:"SoftwareName,omitempty" xml:"SoftwareName,omitempty"`
	SoftwareVersions []*string `json:"SoftwareVersions,omitempty" xml:"SoftwareVersions,omitempty" type:"Repeated"`
}

func (s ListAuthorizedSoftwareResponseBodySoftwares) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizedSoftwareResponseBodySoftwares) GoString() string {
	return s.String()
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetHelpLink(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.HelpLink = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetLastModified(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.LastModified = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetLocations(v []*string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.Locations = v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetPromotion(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.Promotion = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareDefaultVersion(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareDefaultVersion = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareDescription(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareDescription = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareIcon(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareIcon = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareLicenseFee(v float32) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareLicenseFee = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareLongName(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareLongName = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareName(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareName = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareVersions(v []*string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareVersions = v
	return s
}

type ListAuthorizedSoftwareResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedSoftwareResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizedSoftwareResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedSoftwareResponse) SetHeaders(v map[string]*string) *ListAuthorizedSoftwareResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedSoftwareResponse) SetStatusCode(v int32) *ListAuthorizedSoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedSoftwareResponse) SetBody(v *ListAuthorizedSoftwareResponseBody) *ListAuthorizedSoftwareResponse {
	s.Body = v
	return s
}

type ListContainerImagesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0f2cc944-67db-4fb1-bb10-f136cda74bda
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListContainerImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesRequest) GoString() string {
	return s.String()
}

func (s *ListContainerImagesRequest) SetLocation(v string) *ListContainerImagesRequest {
	s.Location = &v
	return s
}

func (s *ListContainerImagesRequest) SetMaxResults(v int32) *ListContainerImagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListContainerImagesRequest) SetNextToken(v string) *ListContainerImagesRequest {
	s.NextToken = &v
	return s
}

type ListContainerImagesResponseBody struct {
	ContainerImages []*ListContainerImagesResponseBodyContainerImages `json:"ContainerImages,omitempty" xml:"ContainerImages,omitempty" type:"Repeated"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ba6d88e2-a6f2-4d2c-a2ae-68b76b078445
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListContainerImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponseBody) SetContainerImages(v []*ListContainerImagesResponseBodyContainerImages) *ListContainerImagesResponseBody {
	s.ContainerImages = v
	return s
}

func (s *ListContainerImagesResponseBody) SetHostId(v string) *ListContainerImagesResponseBody {
	s.HostId = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetMaxResults(v int32) *ListContainerImagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetNextToken(v string) *ListContainerImagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetRequestId(v string) *ListContainerImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetTotalCount(v int32) *ListContainerImagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListContainerImagesResponseBodyContainerImages struct {
	// example:
	//
	// a suite of programs for interacting with high-throughput sequencing data.
	ContainerImageDescription *string `json:"ContainerImageDescription,omitempty" xml:"ContainerImageDescription,omitempty"`
	// example:
	//
	// samtools
	ContainerImageName *string `json:"ContainerImageName,omitempty" xml:"ContainerImageName,omitempty"`
	// example:
	//
	// samtools
	ContainerImageNamespace *string   `json:"ContainerImageNamespace,omitempty" xml:"ContainerImageNamespace,omitempty"`
	ContainerImageVersions  []*string `json:"ContainerImageVersions,omitempty" xml:"ContainerImageVersions,omitempty" type:"Repeated"`
	// example:
	//
	// registry-vpc.cn-shenzhen.aliyuncs.com
	ContainerRegistry *string `json:"ContainerRegistry,omitempty" xml:"ContainerRegistry,omitempty"`
	// example:
	//
	// yyyy/MM/dd HH:mm:ss
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// example:
	//
	// cn-shenzhen
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s ListContainerImagesResponseBodyContainerImages) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponseBodyContainerImages) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponseBodyContainerImages) SetContainerImageDescription(v string) *ListContainerImagesResponseBodyContainerImages {
	s.ContainerImageDescription = &v
	return s
}

func (s *ListContainerImagesResponseBodyContainerImages) SetContainerImageName(v string) *ListContainerImagesResponseBodyContainerImages {
	s.ContainerImageName = &v
	return s
}

func (s *ListContainerImagesResponseBodyContainerImages) SetContainerImageNamespace(v string) *ListContainerImagesResponseBodyContainerImages {
	s.ContainerImageNamespace = &v
	return s
}

func (s *ListContainerImagesResponseBodyContainerImages) SetContainerImageVersions(v []*string) *ListContainerImagesResponseBodyContainerImages {
	s.ContainerImageVersions = v
	return s
}

func (s *ListContainerImagesResponseBodyContainerImages) SetContainerRegistry(v string) *ListContainerImagesResponseBodyContainerImages {
	s.ContainerRegistry = &v
	return s
}

func (s *ListContainerImagesResponseBodyContainerImages) SetLastModified(v string) *ListContainerImagesResponseBodyContainerImages {
	s.LastModified = &v
	return s
}

func (s *ListContainerImagesResponseBodyContainerImages) SetLocation(v string) *ListContainerImagesResponseBodyContainerImages {
	s.Location = &v
	return s
}

type ListContainerImagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContainerImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContainerImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponse) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponse) SetHeaders(v map[string]*string) *ListContainerImagesResponse {
	s.Headers = v
	return s
}

func (s *ListContainerImagesResponse) SetStatusCode(v int32) *ListContainerImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContainerImagesResponse) SetBody(v *ListContainerImagesResponseBody) *ListContainerImagesResponse {
	s.Body = v
	return s
}

type ListEntitiesRequest struct {
	// example:
	//
	// false
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MTIzNDU2P1dvcmtzcGFjZT93cy1mMDE4OTgxNGU1ZTkK
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// createtime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ListEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListEntitiesRequest) SetIsReversed(v bool) *ListEntitiesRequest {
	s.IsReversed = &v
	return s
}

func (s *ListEntitiesRequest) SetMaxResults(v int32) *ListEntitiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEntitiesRequest) SetNextToken(v string) *ListEntitiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListEntitiesRequest) SetOrderBy(v string) *ListEntitiesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListEntitiesRequest) SetWorkspace(v string) *ListEntitiesRequest {
	s.Workspace = &v
	return s
}

type ListEntitiesResponseBody struct {
	Entities []*ListEntitiesResponseBodyEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MTIzNDU2P1dvcmtzcGFjZT93cy1mMDE4OTgxNGU1ZTkK
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBody) SetEntities(v []*ListEntitiesResponseBodyEntities) *ListEntitiesResponseBody {
	s.Entities = v
	return s
}

func (s *ListEntitiesResponseBody) SetHostId(v string) *ListEntitiesResponseBody {
	s.HostId = &v
	return s
}

func (s *ListEntitiesResponseBody) SetMaxResults(v int32) *ListEntitiesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEntitiesResponseBody) SetNextToken(v string) *ListEntitiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEntitiesResponseBody) SetRequestId(v string) *ListEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEntitiesResponseBody) SetTotalCount(v int32) *ListEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

type ListEntitiesResponseBodyEntities struct {
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s ListEntitiesResponseBodyEntities) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesResponseBodyEntities) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBodyEntities) SetEntityType(v string) *ListEntitiesResponseBodyEntities {
	s.EntityType = &v
	return s
}

type ListEntitiesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponse) SetHeaders(v map[string]*string) *ListEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListEntitiesResponse) SetStatusCode(v int32) *ListEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEntitiesResponse) SetBody(v *ListEntitiesResponseBody) *ListEntitiesResponse {
	s.Body = v
	return s
}

type ListEntityItemsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// false
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MTIzNDU2P1dvcmtzcGFjZT93cy1mMDE4OTgxNGU1ZTkK
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// createtime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// test
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ListEntityItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEntityItemsRequest) GoString() string {
	return s.String()
}

func (s *ListEntityItemsRequest) SetEntityType(v string) *ListEntityItemsRequest {
	s.EntityType = &v
	return s
}

func (s *ListEntityItemsRequest) SetIsReversed(v bool) *ListEntityItemsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListEntityItemsRequest) SetMaxResults(v int32) *ListEntityItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEntityItemsRequest) SetNextToken(v string) *ListEntityItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEntityItemsRequest) SetOrderBy(v string) *ListEntityItemsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListEntityItemsRequest) SetSearch(v string) *ListEntityItemsRequest {
	s.Search = &v
	return s
}

func (s *ListEntityItemsRequest) SetWorkspace(v string) *ListEntityItemsRequest {
	s.Workspace = &v
	return s
}

type ListEntityItemsResponseBody struct {
	EntityItems []*ListEntityItemsResponseBodyEntityItems `json:"EntityItems,omitempty" xml:"EntityItems,omitempty" type:"Repeated"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MTIzNDU2P1dvcmtzcGFjZT93cy1mMDE4OTgxNGU1ZTkK
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEntityItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEntityItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEntityItemsResponseBody) SetEntityItems(v []*ListEntityItemsResponseBodyEntityItems) *ListEntityItemsResponseBody {
	s.EntityItems = v
	return s
}

func (s *ListEntityItemsResponseBody) SetHostId(v string) *ListEntityItemsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListEntityItemsResponseBody) SetMaxResults(v int32) *ListEntityItemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEntityItemsResponseBody) SetNextToken(v string) *ListEntityItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEntityItemsResponseBody) SetRequestId(v string) *ListEntityItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEntityItemsResponseBody) SetTotalCount(v int32) *ListEntityItemsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEntityItemsResponseBodyEntityItems struct {
	EntityData map[string]*string `json:"EntityData,omitempty" xml:"EntityData,omitempty"`
	// example:
	//
	// s1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
}

func (s ListEntityItemsResponseBodyEntityItems) String() string {
	return tea.Prettify(s)
}

func (s ListEntityItemsResponseBodyEntityItems) GoString() string {
	return s.String()
}

func (s *ListEntityItemsResponseBodyEntityItems) SetEntityData(v map[string]*string) *ListEntityItemsResponseBodyEntityItems {
	s.EntityData = v
	return s
}

func (s *ListEntityItemsResponseBodyEntityItems) SetEntityName(v string) *ListEntityItemsResponseBodyEntityItems {
	s.EntityName = &v
	return s
}

type ListEntityItemsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEntityItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEntityItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEntityItemsResponse) GoString() string {
	return s.String()
}

func (s *ListEntityItemsResponse) SetHeaders(v map[string]*string) *ListEntityItemsResponse {
	s.Headers = v
	return s
}

func (s *ListEntityItemsResponse) SetStatusCode(v int32) *ListEntityItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEntityItemsResponse) SetBody(v *ListEntityItemsResponseBody) *ListEntityItemsResponse {
	s.Body = v
	return s
}

type ListGlobalAppsRequest struct {
	// example:
	//
	// OpenAccess
	AppScope *string `json:"AppScope,omitempty" xml:"AppScope,omitempty"`
	// example:
	//
	// 群体变异
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// true
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// example:
	//
	// cn-shenzhen
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0f2cc944-67db-4fb1-bb10-f136cda74bda
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// AppName, LastModified, Used
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 群体变异
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// example:
	//
	// Sentieon
	Toolkit *string `json:"Toolkit,omitempty" xml:"Toolkit,omitempty"`
}

func (s ListGlobalAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalAppsRequest) GoString() string {
	return s.String()
}

func (s *ListGlobalAppsRequest) SetAppScope(v string) *ListGlobalAppsRequest {
	s.AppScope = &v
	return s
}

func (s *ListGlobalAppsRequest) SetCategory(v string) *ListGlobalAppsRequest {
	s.Category = &v
	return s
}

func (s *ListGlobalAppsRequest) SetIsReversed(v bool) *ListGlobalAppsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListGlobalAppsRequest) SetLocation(v string) *ListGlobalAppsRequest {
	s.Location = &v
	return s
}

func (s *ListGlobalAppsRequest) SetMaxResults(v int32) *ListGlobalAppsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGlobalAppsRequest) SetNextToken(v string) *ListGlobalAppsRequest {
	s.NextToken = &v
	return s
}

func (s *ListGlobalAppsRequest) SetOrderBy(v string) *ListGlobalAppsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListGlobalAppsRequest) SetSearch(v string) *ListGlobalAppsRequest {
	s.Search = &v
	return s
}

func (s *ListGlobalAppsRequest) SetToolkit(v string) *ListGlobalAppsRequest {
	s.Toolkit = &v
	return s
}

type ListGlobalAppsResponseBody struct {
	GlobalApps []*ListGlobalAppsResponseBodyGlobalApps `json:"GlobalApps,omitempty" xml:"GlobalApps,omitempty" type:"Repeated"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ba6d88e2-a6f2-4d2c-a2ae-68b76b078445
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGlobalAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGlobalAppsResponseBody) SetGlobalApps(v []*ListGlobalAppsResponseBodyGlobalApps) *ListGlobalAppsResponseBody {
	s.GlobalApps = v
	return s
}

func (s *ListGlobalAppsResponseBody) SetHostId(v string) *ListGlobalAppsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetMaxResults(v int32) *ListGlobalAppsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetNextToken(v string) *ListGlobalAppsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetRequestId(v string) *ListGlobalAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetTotalCount(v int64) *ListGlobalAppsResponseBody {
	s.TotalCount = &v
	return s
}

type ListGlobalAppsResponseBodyGlobalApps struct {
	// example:
	//
	// 1.0
	AppDefaultVersion *string `json:"AppDefaultVersion,omitempty" xml:"AppDefaultVersion,omitempty"`
	// example:
	//
	// This is app description.
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// example:
	//
	// 实际使用时，计算该应用消耗的资源费用。
	AppFee *string `json:"AppFee,omitempty" xml:"AppFee,omitempty"`
	// example:
	//
	// joint-calling
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// OpenAccess
	AppScope   *string   `json:"AppScope,omitempty" xml:"AppScope,omitempty"`
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// example:
	//
	// yyyy/MM/dd HH:mm:ss
	LastModified *string   `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Locations    []*string `json:"Locations,omitempty" xml:"Locations,omitempty" type:"Repeated"`
	// example:
	//
	// sentieon
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// example:
	//
	// true
	Pinned *bool `json:"Pinned,omitempty" xml:"Pinned,omitempty"`
	// example:
	//
	// Sentieon
	Toolkit *string `json:"Toolkit,omitempty" xml:"Toolkit,omitempty"`
}

func (s ListGlobalAppsResponseBodyGlobalApps) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalAppsResponseBodyGlobalApps) GoString() string {
	return s.String()
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetAppDefaultVersion(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.AppDefaultVersion = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetAppDescription(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.AppDescription = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetAppFee(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.AppFee = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetAppName(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.AppName = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetAppScope(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.AppScope = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetCategories(v []*string) *ListGlobalAppsResponseBodyGlobalApps {
	s.Categories = v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetLastModified(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.LastModified = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetLocations(v []*string) *ListGlobalAppsResponseBodyGlobalApps {
	s.Locations = v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetNamespaceName(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.NamespaceName = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetPinned(v bool) *ListGlobalAppsResponseBodyGlobalApps {
	s.Pinned = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetToolkit(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.Toolkit = &v
	return s
}

type ListGlobalAppsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGlobalAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGlobalAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalAppsResponse) GoString() string {
	return s.String()
}

func (s *ListGlobalAppsResponse) SetHeaders(v map[string]*string) *ListGlobalAppsResponse {
	s.Headers = v
	return s
}

func (s *ListGlobalAppsResponse) SetStatusCode(v int32) *ListGlobalAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGlobalAppsResponse) SetBody(v *ListGlobalAppsResponseBody) *ListGlobalAppsResponse {
	s.Body = v
	return s
}

type ListPublicDatasetRequest struct {
	// example:
	//
	// true
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0f2cc944-67db-4fb1-bb10-f136cda74bda
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// LastModified, DatasetName
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 西蒙斯基因组多样性计划
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// example:
	//
	// 遗传多样性
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListPublicDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetRequest) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetRequest) SetIsReversed(v bool) *ListPublicDatasetRequest {
	s.IsReversed = &v
	return s
}

func (s *ListPublicDatasetRequest) SetMaxResults(v int32) *ListPublicDatasetRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicDatasetRequest) SetNextToken(v string) *ListPublicDatasetRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetRequest) SetOrderBy(v string) *ListPublicDatasetRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPublicDatasetRequest) SetSearch(v string) *ListPublicDatasetRequest {
	s.Search = &v
	return s
}

func (s *ListPublicDatasetRequest) SetTag(v string) *ListPublicDatasetRequest {
	s.Tag = &v
	return s
}

type ListPublicDatasetResponseBody struct {
	Datasets []*ListPublicDatasetResponseBodyDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ba6d88e2-a6f2-4d2c-a2ae-68b76b078445
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublicDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetResponseBody) SetDatasets(v []*ListPublicDatasetResponseBodyDatasets) *ListPublicDatasetResponseBody {
	s.Datasets = v
	return s
}

func (s *ListPublicDatasetResponseBody) SetHostId(v string) *ListPublicDatasetResponseBody {
	s.HostId = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetMaxResults(v int32) *ListPublicDatasetResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetNextToken(v string) *ListPublicDatasetResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetRequestId(v string) *ListPublicDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetTotalCount(v int32) *ListPublicDatasetResponseBody {
	s.TotalCount = &v
	return s
}

type ListPublicDatasetResponseBodyDatasets struct {
	// example:
	//
	// This is description about dataset.
	About *string `json:"About,omitempty" xml:"About,omitempty"`
	// example:
	//
	// This are access requirements.
	AccessRequirements *string `json:"AccessRequirements,omitempty" xml:"AccessRequirements,omitempty"`
	// example:
	//
	// This is copyright.
	Copyright *string `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	// example:
	//
	// This are access requirements
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	DatasetId          *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 西蒙斯基因组多样性计划
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// yyyy/MM/dd HH:mm:ss
	LastModified *string   `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Locations    []*string `json:"Locations,omitempty" xml:"Locations,omitempty" type:"Repeated"`
	Tags         []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 源站更新后及时同步
	UpdateFrequency *string `json:"UpdateFrequency,omitempty" xml:"UpdateFrequency,omitempty"`
}

func (s ListPublicDatasetResponseBodyDatasets) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetResponseBodyDatasets) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetResponseBodyDatasets) SetAbout(v string) *ListPublicDatasetResponseBodyDatasets {
	s.About = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetAccessRequirements(v string) *ListPublicDatasetResponseBodyDatasets {
	s.AccessRequirements = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetCopyright(v string) *ListPublicDatasetResponseBodyDatasets {
	s.Copyright = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetDatasetDescription(v string) *ListPublicDatasetResponseBodyDatasets {
	s.DatasetDescription = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetDatasetId(v string) *ListPublicDatasetResponseBodyDatasets {
	s.DatasetId = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetDatasetName(v string) *ListPublicDatasetResponseBodyDatasets {
	s.DatasetName = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetLastModified(v string) *ListPublicDatasetResponseBodyDatasets {
	s.LastModified = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetLocations(v []*string) *ListPublicDatasetResponseBodyDatasets {
	s.Locations = v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetTags(v []*string) *ListPublicDatasetResponseBodyDatasets {
	s.Tags = v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetUpdateFrequency(v string) *ListPublicDatasetResponseBodyDatasets {
	s.UpdateFrequency = &v
	return s
}

type ListPublicDatasetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetResponse) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetResponse) SetHeaders(v map[string]*string) *ListPublicDatasetResponse {
	s.Headers = v
	return s
}

func (s *ListPublicDatasetResponse) SetStatusCode(v int32) *ListPublicDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicDatasetResponse) SetBody(v *ListPublicDatasetResponseBody) *ListPublicDatasetResponse {
	s.Body = v
	return s
}

type ListPublicDatasetEntitiesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 西蒙斯基因组多样性计划
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// true
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// example:
	//
	// cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0f2cc944-67db-4fb1-bb10-f136cda74bda
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// LastModified
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
}

func (s ListPublicDatasetEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntitiesRequest) SetDatasetName(v string) *ListPublicDatasetEntitiesRequest {
	s.DatasetName = &v
	return s
}

func (s *ListPublicDatasetEntitiesRequest) SetIsReversed(v bool) *ListPublicDatasetEntitiesRequest {
	s.IsReversed = &v
	return s
}

func (s *ListPublicDatasetEntitiesRequest) SetLocation(v string) *ListPublicDatasetEntitiesRequest {
	s.Location = &v
	return s
}

func (s *ListPublicDatasetEntitiesRequest) SetMaxResults(v int32) *ListPublicDatasetEntitiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicDatasetEntitiesRequest) SetNextToken(v string) *ListPublicDatasetEntitiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetEntitiesRequest) SetOrderBy(v string) *ListPublicDatasetEntitiesRequest {
	s.OrderBy = &v
	return s
}

type ListPublicDatasetEntitiesResponseBody struct {
	// example:
	//
	// 西蒙斯基因组多样性计划
	DatasetName *string                                          `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Entities    []*ListPublicDatasetEntitiesResponseBodyEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ba6d88e2-a6f2-4d2c-a2ae-68b76b078445
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublicDatasetEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntitiesResponseBody) SetDatasetName(v string) *ListPublicDatasetEntitiesResponseBody {
	s.DatasetName = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetEntities(v []*ListPublicDatasetEntitiesResponseBodyEntities) *ListPublicDatasetEntitiesResponseBody {
	s.Entities = v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetHostId(v string) *ListPublicDatasetEntitiesResponseBody {
	s.HostId = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetMaxResults(v int32) *ListPublicDatasetEntitiesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetNextToken(v string) *ListPublicDatasetEntitiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetRequestId(v string) *ListPublicDatasetEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetTotalCount(v int32) *ListPublicDatasetEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

type ListPublicDatasetEntitiesResponseBodyEntities struct {
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s ListPublicDatasetEntitiesResponseBodyEntities) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntitiesResponseBodyEntities) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntitiesResponseBodyEntities) SetEntityType(v string) *ListPublicDatasetEntitiesResponseBodyEntities {
	s.EntityType = &v
	return s
}

type ListPublicDatasetEntitiesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicDatasetEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicDatasetEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntitiesResponse) SetHeaders(v map[string]*string) *ListPublicDatasetEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListPublicDatasetEntitiesResponse) SetStatusCode(v int32) *ListPublicDatasetEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponse) SetBody(v *ListPublicDatasetEntitiesResponseBody) *ListPublicDatasetEntitiesResponse {
	s.Body = v
	return s
}

type ListPublicDatasetEntityItemsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 西蒙斯基因组多样性计划
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// true
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// example:
	//
	// cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0f2cc944-67db-4fb1-bb10-f136cda74bda
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// LastModified
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// s1
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
}

func (s ListPublicDatasetEntityItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntityItemsRequest) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntityItemsRequest) SetDatasetName(v string) *ListPublicDatasetEntityItemsRequest {
	s.DatasetName = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetEntityType(v string) *ListPublicDatasetEntityItemsRequest {
	s.EntityType = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetIsReversed(v bool) *ListPublicDatasetEntityItemsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetLocation(v string) *ListPublicDatasetEntityItemsRequest {
	s.Location = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetMaxResults(v int32) *ListPublicDatasetEntityItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetNextToken(v string) *ListPublicDatasetEntityItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetOrderBy(v string) *ListPublicDatasetEntityItemsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetSearch(v string) *ListPublicDatasetEntityItemsRequest {
	s.Search = &v
	return s
}

type ListPublicDatasetEntityItemsResponseBody struct {
	// example:
	//
	// 西蒙斯基因组多样性计划
	DatasetName *string                                                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	EntityItems []*ListPublicDatasetEntityItemsResponseBodyEntityItems `json:"EntityItems,omitempty" xml:"EntityItems,omitempty" type:"Repeated"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ba6d88e2-a6f2-4d2c-a2ae-68b76b078445
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublicDatasetEntityItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntityItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetDatasetName(v string) *ListPublicDatasetEntityItemsResponseBody {
	s.DatasetName = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetEntityItems(v []*ListPublicDatasetEntityItemsResponseBodyEntityItems) *ListPublicDatasetEntityItemsResponseBody {
	s.EntityItems = v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetHostId(v string) *ListPublicDatasetEntityItemsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetMaxResults(v int32) *ListPublicDatasetEntityItemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetNextToken(v string) *ListPublicDatasetEntityItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetRequestId(v string) *ListPublicDatasetEntityItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetTotalCount(v int32) *ListPublicDatasetEntityItemsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPublicDatasetEntityItemsResponseBodyEntityItems struct {
	EntityData map[string]*string `json:"EntityData,omitempty" xml:"EntityData,omitempty"`
	// example:
	//
	// sample
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
}

func (s ListPublicDatasetEntityItemsResponseBodyEntityItems) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntityItemsResponseBodyEntityItems) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntityItemsResponseBodyEntityItems) SetEntityData(v map[string]*string) *ListPublicDatasetEntityItemsResponseBodyEntityItems {
	s.EntityData = v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBodyEntityItems) SetEntityName(v string) *ListPublicDatasetEntityItemsResponseBodyEntityItems {
	s.EntityName = &v
	return s
}

type ListPublicDatasetEntityItemsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicDatasetEntityItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicDatasetEntityItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntityItemsResponse) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntityItemsResponse) SetHeaders(v map[string]*string) *ListPublicDatasetEntityItemsResponse {
	s.Headers = v
	return s
}

func (s *ListPublicDatasetEntityItemsResponse) SetStatusCode(v int32) *ListPublicDatasetEntityItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponse) SetBody(v *ListPublicDatasetEntityItemsResponseBody) *ListPublicDatasetEntityItemsResponse {
	s.Body = v
	return s
}

type ListPublicDatasetTagsRequest struct {
	// example:
	//
	// true
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0f2cc944-67db-4fb1-bb10-f136cda74bda
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// LastModified
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 遗传多样性
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
}

func (s ListPublicDatasetTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetTagsRequest) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetTagsRequest) SetIsReversed(v bool) *ListPublicDatasetTagsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListPublicDatasetTagsRequest) SetMaxResults(v int32) *ListPublicDatasetTagsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicDatasetTagsRequest) SetNextToken(v string) *ListPublicDatasetTagsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetTagsRequest) SetOrderBy(v string) *ListPublicDatasetTagsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPublicDatasetTagsRequest) SetSearch(v string) *ListPublicDatasetTagsRequest {
	s.Search = &v
	return s
}

type ListPublicDatasetTagsResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ba6d88e2-a6f2-4d2c-a2ae-68b76b078445
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublicDatasetTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetTagsResponseBody) SetHostId(v string) *ListPublicDatasetTagsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetMaxResults(v int32) *ListPublicDatasetTagsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetNextToken(v string) *ListPublicDatasetTagsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetRequestId(v string) *ListPublicDatasetTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetTags(v []*string) *ListPublicDatasetTagsResponseBody {
	s.Tags = v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetTotalCount(v int32) *ListPublicDatasetTagsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPublicDatasetTagsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicDatasetTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicDatasetTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetTagsResponse) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetTagsResponse) SetHeaders(v map[string]*string) *ListPublicDatasetTagsResponse {
	s.Headers = v
	return s
}

func (s *ListPublicDatasetTagsResponse) SetStatusCode(v int32) *ListPublicDatasetTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicDatasetTagsResponse) SetBody(v *ListPublicDatasetTagsResponseBody) *ListPublicDatasetTagsResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId  *string                           `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Regions []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetHostId(v string) *ListRegionsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegions struct {
	// example:
	//
	// 华北2-北京
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionEndpoint(v string) *ListRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListRunsRequest struct {
	// example:
	//
	// MyApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 4.1.4.1-v1.0
	AppRevision *string `json:"AppRevision,omitempty" xml:"AppRevision,omitempty"`
	// example:
	//
	// S001
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// false
	GetTotal *bool `json:"GetTotal,omitempty" xml:"GetTotal,omitempty"`
	// example:
	//
	// false
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// example:
	//
	// env=test
	LabelSelector *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MTIzNDU2P1dvcmtzcGFjZT93cy1mMDE4OTgxNGU1ZTkK
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// createtime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// run-13BeYuxc0fxO24uA75UppTiMRoQ
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// sub-13BeXWnGvrOl3MC124tazcJVKaV
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ListRunsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRunsRequest) GoString() string {
	return s.String()
}

func (s *ListRunsRequest) SetAppName(v string) *ListRunsRequest {
	s.AppName = &v
	return s
}

func (s *ListRunsRequest) SetAppRevision(v string) *ListRunsRequest {
	s.AppRevision = &v
	return s
}

func (s *ListRunsRequest) SetEntityName(v string) *ListRunsRequest {
	s.EntityName = &v
	return s
}

func (s *ListRunsRequest) SetEntityType(v string) *ListRunsRequest {
	s.EntityType = &v
	return s
}

func (s *ListRunsRequest) SetGetTotal(v bool) *ListRunsRequest {
	s.GetTotal = &v
	return s
}

func (s *ListRunsRequest) SetIsReversed(v bool) *ListRunsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListRunsRequest) SetLabelSelector(v string) *ListRunsRequest {
	s.LabelSelector = &v
	return s
}

func (s *ListRunsRequest) SetMaxResults(v int32) *ListRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRunsRequest) SetNextToken(v string) *ListRunsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRunsRequest) SetOrderBy(v string) *ListRunsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListRunsRequest) SetSearch(v string) *ListRunsRequest {
	s.Search = &v
	return s
}

func (s *ListRunsRequest) SetStatus(v string) *ListRunsRequest {
	s.Status = &v
	return s
}

func (s *ListRunsRequest) SetSubmissionId(v string) *ListRunsRequest {
	s.SubmissionId = &v
	return s
}

func (s *ListRunsRequest) SetWorkspace(v string) *ListRunsRequest {
	s.Workspace = &v
	return s
}

type ListRunsResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MTIzNDU2P1dvcmtzcGFjZT93cy1mMDE4OTgxNGU1ZTkK
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Runs      []*ListRunsResponseBodyRuns `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRunsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRunsResponseBody) SetHostId(v string) *ListRunsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListRunsResponseBody) SetMaxResults(v int32) *ListRunsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRunsResponseBody) SetNextToken(v string) *ListRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRunsResponseBody) SetRequestId(v string) *ListRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRunsResponseBody) SetRuns(v []*ListRunsResponseBodyRuns) *ListRunsResponseBody {
	s.Runs = v
	return s
}

func (s *ListRunsResponseBody) SetTotalCount(v int32) *ListRunsResponseBody {
	s.TotalCount = &v
	return s
}

type ListRunsResponseBodyRuns struct {
	// example:
	//
	// MyApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	AppRevision *string `json:"AppRevision,omitempty" xml:"AppRevision,omitempty"`
	// example:
	//
	// 2021-06-10T13:30:20.414557061Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// {}
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	// example:
	//
	// 2021-06-10T13:38:20.414557061Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// s1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// oss://my-bucket/my-execdir/
	ExecuteDirectory *string                                 `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	ExecuteOptions   *ListRunsResponseBodyRunsExecuteOptions `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty" type:"Struct"`
	// example:
	//
	// {     "workflow1.task1.param1": "helloooo.txt",     "workflow1.task1.param2": "abc",     "workflow1.task2.param1": "new_helloooo.txt",     "workflow1.task2.param2": "def" }
	Inputs *string            `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// run-13BeYuxc0fxO24uA75UppTiMRoQ
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// example:
	//
	// test-run
	RunName *string `json:"RunName,omitempty" xml:"RunName,omitempty"`
	// example:
	//
	// EasyGene
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 2021-06-10T13:31:20.414557061Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// sub-13BeXWnGvrOl3MC124tazcJVKaV
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ListRunsResponseBodyRuns) String() string {
	return tea.Prettify(s)
}

func (s ListRunsResponseBodyRuns) GoString() string {
	return s.String()
}

func (s *ListRunsResponseBodyRuns) SetAppName(v string) *ListRunsResponseBodyRuns {
	s.AppName = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetAppRevision(v string) *ListRunsResponseBodyRuns {
	s.AppRevision = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetCreateTime(v string) *ListRunsResponseBodyRuns {
	s.CreateTime = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetDefaultRuntime(v string) *ListRunsResponseBodyRuns {
	s.DefaultRuntime = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetEndTime(v string) *ListRunsResponseBodyRuns {
	s.EndTime = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetEntityName(v string) *ListRunsResponseBodyRuns {
	s.EntityName = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetEntityType(v string) *ListRunsResponseBodyRuns {
	s.EntityType = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetExecuteDirectory(v string) *ListRunsResponseBodyRuns {
	s.ExecuteDirectory = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetExecuteOptions(v *ListRunsResponseBodyRunsExecuteOptions) *ListRunsResponseBodyRuns {
	s.ExecuteOptions = v
	return s
}

func (s *ListRunsResponseBodyRuns) SetInputs(v string) *ListRunsResponseBodyRuns {
	s.Inputs = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetLabels(v map[string]*string) *ListRunsResponseBodyRuns {
	s.Labels = v
	return s
}

func (s *ListRunsResponseBodyRuns) SetRegionId(v string) *ListRunsResponseBodyRuns {
	s.RegionId = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetRunId(v string) *ListRunsResponseBodyRuns {
	s.RunId = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetRunName(v string) *ListRunsResponseBodyRuns {
	s.RunName = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetSource(v string) *ListRunsResponseBodyRuns {
	s.Source = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetStartTime(v string) *ListRunsResponseBodyRuns {
	s.StartTime = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetStatus(v string) *ListRunsResponseBodyRuns {
	s.Status = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetSubmissionId(v string) *ListRunsResponseBodyRuns {
	s.SubmissionId = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetWorkspace(v string) *ListRunsResponseBodyRuns {
	s.Workspace = &v
	return s
}

type ListRunsResponseBodyRunsExecuteOptions struct {
	// example:
	//
	// false
	CallCaching *bool `json:"CallCaching,omitempty" xml:"CallCaching,omitempty"`
	// example:
	//
	// false
	DeleteIntermediateResults *bool `json:"DeleteIntermediateResults,omitempty" xml:"DeleteIntermediateResults,omitempty"`
	// example:
	//
	// NoNewCalls
	FailureMode *string `json:"FailureMode,omitempty" xml:"FailureMode,omitempty"`
	// example:
	//
	// false
	UseRelativeOutputPaths *bool `json:"UseRelativeOutputPaths,omitempty" xml:"UseRelativeOutputPaths,omitempty"`
}

func (s ListRunsResponseBodyRunsExecuteOptions) String() string {
	return tea.Prettify(s)
}

func (s ListRunsResponseBodyRunsExecuteOptions) GoString() string {
	return s.String()
}

func (s *ListRunsResponseBodyRunsExecuteOptions) SetCallCaching(v bool) *ListRunsResponseBodyRunsExecuteOptions {
	s.CallCaching = &v
	return s
}

func (s *ListRunsResponseBodyRunsExecuteOptions) SetDeleteIntermediateResults(v bool) *ListRunsResponseBodyRunsExecuteOptions {
	s.DeleteIntermediateResults = &v
	return s
}

func (s *ListRunsResponseBodyRunsExecuteOptions) SetFailureMode(v string) *ListRunsResponseBodyRunsExecuteOptions {
	s.FailureMode = &v
	return s
}

func (s *ListRunsResponseBodyRunsExecuteOptions) SetUseRelativeOutputPaths(v bool) *ListRunsResponseBodyRunsExecuteOptions {
	s.UseRelativeOutputPaths = &v
	return s
}

type ListRunsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRunsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRunsResponse) GoString() string {
	return s.String()
}

func (s *ListRunsResponse) SetHeaders(v map[string]*string) *ListRunsResponse {
	s.Headers = v
	return s
}

func (s *ListRunsResponse) SetStatusCode(v int32) *ListRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRunsResponse) SetBody(v *ListRunsResponseBody) *ListRunsResponse {
	s.Body = v
	return s
}

type ListSubmissionsRequest struct {
	// example:
	//
	// true
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Next Token
	//
	// example:
	//
	// TestToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// createtime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// sub-13BeXWnGvrOl3MC124tazcJVKaV
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// example:
	//
	// Failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ListSubmissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubmissionsRequest) GoString() string {
	return s.String()
}

func (s *ListSubmissionsRequest) SetIsReversed(v bool) *ListSubmissionsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListSubmissionsRequest) SetMaxResults(v int32) *ListSubmissionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSubmissionsRequest) SetNextToken(v string) *ListSubmissionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSubmissionsRequest) SetOrderBy(v string) *ListSubmissionsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListSubmissionsRequest) SetSearch(v string) *ListSubmissionsRequest {
	s.Search = &v
	return s
}

func (s *ListSubmissionsRequest) SetStatus(v string) *ListSubmissionsRequest {
	s.Status = &v
	return s
}

func (s *ListSubmissionsRequest) SetWorkspace(v string) *ListSubmissionsRequest {
	s.Workspace = &v
	return s
}

type ListSubmissionsResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// TestToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Submissions []*ListSubmissionsResponseBodySubmissions `json:"Submissions,omitempty" xml:"Submissions,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSubmissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubmissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubmissionsResponseBody) SetHostId(v string) *ListSubmissionsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListSubmissionsResponseBody) SetMaxResults(v int32) *ListSubmissionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSubmissionsResponseBody) SetNextToken(v string) *ListSubmissionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSubmissionsResponseBody) SetRequestId(v string) *ListSubmissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubmissionsResponseBody) SetSubmissions(v []*ListSubmissionsResponseBodySubmissions) *ListSubmissionsResponseBody {
	s.Submissions = v
	return s
}

func (s *ListSubmissionsResponseBody) SetTotalCount(v int32) *ListSubmissionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSubmissionsResponseBodySubmissions struct {
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// TestEntityType
	EntityType *string                                         `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	RunStats   *ListSubmissionsResponseBodySubmissionsRunStats `json:"RunStats,omitempty" xml:"RunStats,omitempty" type:"Struct"`
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// sub-13BeXWnGvrOl3MC124tazcJVKaV
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ListSubmissionsResponseBodySubmissions) String() string {
	return tea.Prettify(s)
}

func (s ListSubmissionsResponseBodySubmissions) GoString() string {
	return s.String()
}

func (s *ListSubmissionsResponseBodySubmissions) SetCreateTime(v string) *ListSubmissionsResponseBodySubmissions {
	s.CreateTime = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetEndTime(v string) *ListSubmissionsResponseBodySubmissions {
	s.EndTime = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetEntityType(v string) *ListSubmissionsResponseBodySubmissions {
	s.EntityType = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetRunStats(v *ListSubmissionsResponseBodySubmissionsRunStats) *ListSubmissionsResponseBodySubmissions {
	s.RunStats = v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetStartTime(v string) *ListSubmissionsResponseBodySubmissions {
	s.StartTime = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetStatus(v string) *ListSubmissionsResponseBodySubmissions {
	s.Status = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetSubmissionId(v string) *ListSubmissionsResponseBodySubmissions {
	s.SubmissionId = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetWorkspace(v string) *ListSubmissionsResponseBodySubmissions {
	s.Workspace = &v
	return s
}

type ListSubmissionsResponseBodySubmissionsRunStats struct {
	// example:
	//
	// 5
	Aborted *int64 `json:"Aborted,omitempty" xml:"Aborted,omitempty"`
	// example:
	//
	// 5
	Aborting *int64 `json:"Aborting,omitempty" xml:"Aborting,omitempty"`
	// example:
	//
	// 5
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// example:
	//
	// 5
	Pending *int64 `json:"Pending,omitempty" xml:"Pending,omitempty"`
	// example:
	//
	// 5
	Running *int64 `json:"Running,omitempty" xml:"Running,omitempty"`
	// example:
	//
	// 5
	Submitted *int64 `json:"Submitted,omitempty" xml:"Submitted,omitempty"`
	// example:
	//
	// 5
	Succeeded *int64 `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s ListSubmissionsResponseBodySubmissionsRunStats) String() string {
	return tea.Prettify(s)
}

func (s ListSubmissionsResponseBodySubmissionsRunStats) GoString() string {
	return s.String()
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetAborted(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Aborted = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetAborting(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Aborting = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetFailed(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Failed = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetPending(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Pending = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetRunning(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Running = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetSubmitted(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Submitted = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetSucceeded(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Succeeded = &v
	return s
}

type ListSubmissionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubmissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubmissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubmissionsResponse) GoString() string {
	return s.String()
}

func (s *ListSubmissionsResponse) SetHeaders(v map[string]*string) *ListSubmissionsResponse {
	s.Headers = v
	return s
}

func (s *ListSubmissionsResponse) SetStatusCode(v int32) *ListSubmissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubmissionsResponse) SetBody(v *ListSubmissionsResponseBody) *ListSubmissionsResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	// example:
	//
	// false
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// example:
	//
	// key1=value1
	LabelSelector *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MTIzNDU2P1dvcmtzcGFjZT93cy1mMDE4OTgxNGU1ZTkK
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// createtime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// test
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetIsReversed(v bool) *ListTemplatesRequest {
	s.IsReversed = &v
	return s
}

func (s *ListTemplatesRequest) SetLabelSelector(v string) *ListTemplatesRequest {
	s.LabelSelector = &v
	return s
}

func (s *ListTemplatesRequest) SetMaxResults(v int32) *ListTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesRequest) SetNextToken(v string) *ListTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesRequest) SetOrderBy(v string) *ListTemplatesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListTemplatesRequest) SetSearch(v string) *ListTemplatesRequest {
	s.Search = &v
	return s
}

func (s *ListTemplatesRequest) SetWorkspace(v string) *ListTemplatesRequest {
	s.Workspace = &v
	return s
}

type ListTemplatesResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MTIzNDU2P1dvcmtzcGFjZT93cy1mMDE4OTgxNGU1ZTkK
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates []*ListTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetHostId(v string) *ListTemplatesResponseBody {
	s.HostId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetMaxResults(v int32) *ListTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesResponseBody) SetNextToken(v string) *ListTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTemplatesResponseBodyTemplates struct {
	// example:
	//
	// my-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	AppRevision *string `json:"AppRevision,omitempty" xml:"AppRevision,omitempty"`
	// example:
	//
	// 2021-06-10T13:30:20.414557061Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// This is a test template
	Description      *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	InputsExpression []*ListTemplatesResponseBodyTemplatesInputsExpression `json:"InputsExpression,omitempty" xml:"InputsExpression,omitempty" type:"Repeated"`
	Labels           map[string]*string                                    `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// 2021-06-10T13:30:20.414557061Z
	LastModifiedTime  *string                                                `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	OutputsExpression []*ListTemplatesResponseBodyTemplatesOutputsExpression `json:"OutputsExpression,omitempty" xml:"OutputsExpression,omitempty" type:"Repeated"`
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// example:
	//
	// test-entity
	RootEntity *string `json:"RootEntity,omitempty" xml:"RootEntity,omitempty"`
	// example:
	//
	// test-template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ListTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplates) SetAppName(v string) *ListTemplatesResponseBodyTemplates {
	s.AppName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetAppRevision(v string) *ListTemplatesResponseBodyTemplates {
	s.AppRevision = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreateTime(v string) *ListTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetDescription(v string) *ListTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetInputsExpression(v []*ListTemplatesResponseBodyTemplatesInputsExpression) *ListTemplatesResponseBodyTemplates {
	s.InputsExpression = v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetLabels(v map[string]*string) *ListTemplatesResponseBodyTemplates {
	s.Labels = v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetLastModifiedTime(v string) *ListTemplatesResponseBodyTemplates {
	s.LastModifiedTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetOutputsExpression(v []*ListTemplatesResponseBodyTemplatesOutputsExpression) *ListTemplatesResponseBodyTemplates {
	s.OutputsExpression = v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetRevisionTag(v string) *ListTemplatesResponseBodyTemplates {
	s.RevisionTag = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetRootEntity(v string) *ListTemplatesResponseBodyTemplates {
	s.RootEntity = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateName(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetWorkspace(v string) *ListTemplatesResponseBodyTemplates {
	s.Workspace = &v
	return s
}

type ListTemplatesResponseBodyTemplatesInputsExpression struct {
	// example:
	//
	// help test
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// 0
	StepOrder *int64 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	// example:
	//
	// task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// fastq1
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// example:
	//
	// File
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// example:
	//
	// oss://bucket/file1.fastq
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s ListTemplatesResponseBodyTemplatesInputsExpression) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplatesInputsExpression) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplatesInputsExpression) SetHelp(v string) *ListTemplatesResponseBodyTemplatesInputsExpression {
	s.Help = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesInputsExpression) SetRequired(v bool) *ListTemplatesResponseBodyTemplatesInputsExpression {
	s.Required = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesInputsExpression) SetStepOrder(v int64) *ListTemplatesResponseBodyTemplatesInputsExpression {
	s.StepOrder = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesInputsExpression) SetTaskName(v string) *ListTemplatesResponseBodyTemplatesInputsExpression {
	s.TaskName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesInputsExpression) SetVariableName(v string) *ListTemplatesResponseBodyTemplatesInputsExpression {
	s.VariableName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesInputsExpression) SetVariableType(v string) *ListTemplatesResponseBodyTemplatesInputsExpression {
	s.VariableType = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesInputsExpression) SetVariableValue(v string) *ListTemplatesResponseBodyTemplatesInputsExpression {
	s.VariableValue = &v
	return s
}

type ListTemplatesResponseBodyTemplatesOutputsExpression struct {
	// example:
	//
	// help test
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// 1
	StepOrder *int64 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	// example:
	//
	// task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// v1
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// example:
	//
	// File
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// example:
	//
	// oss://bucket/test.vcf
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s ListTemplatesResponseBodyTemplatesOutputsExpression) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplatesOutputsExpression) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplatesOutputsExpression) SetHelp(v string) *ListTemplatesResponseBodyTemplatesOutputsExpression {
	s.Help = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesOutputsExpression) SetRequired(v bool) *ListTemplatesResponseBodyTemplatesOutputsExpression {
	s.Required = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesOutputsExpression) SetStepOrder(v int64) *ListTemplatesResponseBodyTemplatesOutputsExpression {
	s.StepOrder = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesOutputsExpression) SetTaskName(v string) *ListTemplatesResponseBodyTemplatesOutputsExpression {
	s.TaskName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesOutputsExpression) SetVariableName(v string) *ListTemplatesResponseBodyTemplatesOutputsExpression {
	s.VariableName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesOutputsExpression) SetVariableType(v string) *ListTemplatesResponseBodyTemplatesOutputsExpression {
	s.VariableType = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesOutputsExpression) SetVariableValue(v string) *ListTemplatesResponseBodyTemplatesOutputsExpression {
	s.VariableValue = &v
	return s
}

type ListTemplatesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) SetHeaders(v map[string]*string) *ListTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesResponse) SetStatusCode(v int32) *ListTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResponseBody) *ListTemplatesResponse {
	s.Body = v
	return s
}

type ListUserActiveRunsRequest struct {
	// example:
	//
	// 3
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListUserActiveRunsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserActiveRunsRequest) GoString() string {
	return s.String()
}

func (s *ListUserActiveRunsRequest) SetMaxResults(v int32) *ListUserActiveRunsRequest {
	s.MaxResults = &v
	return s
}

type ListUserActiveRunsResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MTIzNDU2P1dvcmtzcGFjZT93cy1mMDE4OTgxNGU1ZTkK
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Runs      []*ListUserActiveRunsResponseBodyRuns `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserActiveRunsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserActiveRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserActiveRunsResponseBody) SetHostId(v string) *ListUserActiveRunsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListUserActiveRunsResponseBody) SetMaxResults(v int32) *ListUserActiveRunsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserActiveRunsResponseBody) SetNextToken(v string) *ListUserActiveRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserActiveRunsResponseBody) SetRequestId(v string) *ListUserActiveRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserActiveRunsResponseBody) SetRuns(v []*ListUserActiveRunsResponseBodyRuns) *ListUserActiveRunsResponseBody {
	s.Runs = v
	return s
}

func (s *ListUserActiveRunsResponseBody) SetTotalCount(v int32) *ListUserActiveRunsResponseBody {
	s.TotalCount = &v
	return s
}

type ListUserActiveRunsResponseBodyRuns struct {
	// example:
	//
	// MyApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	AppRevision *string `json:"AppRevision,omitempty" xml:"AppRevision,omitempty"`
	// example:
	//
	// 2021-06-10T13:30:20.414557061Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// {}
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	// example:
	//
	// 2021-06-10T13:30:20.414557061Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// s1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// example:
	//
	// sample
	EntityType       *string                                           `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	ExecuteDirectory *string                                           `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	ExecuteOptions   *ListUserActiveRunsResponseBodyRunsExecuteOptions `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty" type:"Struct"`
	// example:
	//
	// { "workflow1.task1.param1": "helloooo.txt", "workflow1.task1.param2": "abc", "workflow1.task2.param1": "new_helloooo.txt", "workflow1.task2.param2": "def" }
	Inputs *string            `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// run-13BeYuxc0fxO24uA75UppTiMRoQ
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// example:
	//
	// test-run
	RunName *string `json:"RunName,omitempty" xml:"RunName,omitempty"`
	// example:
	//
	// EasyGene
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 2021-06-10T13:30:20.414557061Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// sub-13BeXWnGvrOl3MC124tazcJVKaV
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ListUserActiveRunsResponseBodyRuns) String() string {
	return tea.Prettify(s)
}

func (s ListUserActiveRunsResponseBodyRuns) GoString() string {
	return s.String()
}

func (s *ListUserActiveRunsResponseBodyRuns) SetAppName(v string) *ListUserActiveRunsResponseBodyRuns {
	s.AppName = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetAppRevision(v string) *ListUserActiveRunsResponseBodyRuns {
	s.AppRevision = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetCreateTime(v string) *ListUserActiveRunsResponseBodyRuns {
	s.CreateTime = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetDefaultRuntime(v string) *ListUserActiveRunsResponseBodyRuns {
	s.DefaultRuntime = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetEndTime(v string) *ListUserActiveRunsResponseBodyRuns {
	s.EndTime = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetEntityName(v string) *ListUserActiveRunsResponseBodyRuns {
	s.EntityName = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetEntityType(v string) *ListUserActiveRunsResponseBodyRuns {
	s.EntityType = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetExecuteDirectory(v string) *ListUserActiveRunsResponseBodyRuns {
	s.ExecuteDirectory = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetExecuteOptions(v *ListUserActiveRunsResponseBodyRunsExecuteOptions) *ListUserActiveRunsResponseBodyRuns {
	s.ExecuteOptions = v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetInputs(v string) *ListUserActiveRunsResponseBodyRuns {
	s.Inputs = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetLabels(v map[string]*string) *ListUserActiveRunsResponseBodyRuns {
	s.Labels = v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetRegionId(v string) *ListUserActiveRunsResponseBodyRuns {
	s.RegionId = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetRunId(v string) *ListUserActiveRunsResponseBodyRuns {
	s.RunId = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetRunName(v string) *ListUserActiveRunsResponseBodyRuns {
	s.RunName = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetSource(v string) *ListUserActiveRunsResponseBodyRuns {
	s.Source = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetStartTime(v string) *ListUserActiveRunsResponseBodyRuns {
	s.StartTime = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetStatus(v string) *ListUserActiveRunsResponseBodyRuns {
	s.Status = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetSubmissionId(v string) *ListUserActiveRunsResponseBodyRuns {
	s.SubmissionId = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetWorkspace(v string) *ListUserActiveRunsResponseBodyRuns {
	s.Workspace = &v
	return s
}

type ListUserActiveRunsResponseBodyRunsExecuteOptions struct {
	// example:
	//
	// false
	CallCaching *bool `json:"CallCaching,omitempty" xml:"CallCaching,omitempty"`
	// example:
	//
	// false
	DeleteIntermediateResults *bool `json:"DeleteIntermediateResults,omitempty" xml:"DeleteIntermediateResults,omitempty"`
	// example:
	//
	// NoNewCalls
	FailureMode *string `json:"FailureMode,omitempty" xml:"FailureMode,omitempty"`
	// example:
	//
	// false
	UseRelativeOutputPaths *bool `json:"UseRelativeOutputPaths,omitempty" xml:"UseRelativeOutputPaths,omitempty"`
}

func (s ListUserActiveRunsResponseBodyRunsExecuteOptions) String() string {
	return tea.Prettify(s)
}

func (s ListUserActiveRunsResponseBodyRunsExecuteOptions) GoString() string {
	return s.String()
}

func (s *ListUserActiveRunsResponseBodyRunsExecuteOptions) SetCallCaching(v bool) *ListUserActiveRunsResponseBodyRunsExecuteOptions {
	s.CallCaching = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRunsExecuteOptions) SetDeleteIntermediateResults(v bool) *ListUserActiveRunsResponseBodyRunsExecuteOptions {
	s.DeleteIntermediateResults = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRunsExecuteOptions) SetFailureMode(v string) *ListUserActiveRunsResponseBodyRunsExecuteOptions {
	s.FailureMode = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRunsExecuteOptions) SetUseRelativeOutputPaths(v bool) *ListUserActiveRunsResponseBodyRunsExecuteOptions {
	s.UseRelativeOutputPaths = &v
	return s
}

type ListUserActiveRunsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserActiveRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserActiveRunsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserActiveRunsResponse) GoString() string {
	return s.String()
}

func (s *ListUserActiveRunsResponse) SetHeaders(v map[string]*string) *ListUserActiveRunsResponse {
	s.Headers = v
	return s
}

func (s *ListUserActiveRunsResponse) SetStatusCode(v int32) *ListUserActiveRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserActiveRunsResponse) SetBody(v *ListUserActiveRunsResponseBody) *ListUserActiveRunsResponse {
	s.Body = v
	return s
}

type ListWorkspacesRequest struct {
	// example:
	//
	// true
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// example:
	//
	// env=test
	LabelSelector *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// NextToken
	//
	// example:
	//
	// TestToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// createtime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// TestName
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) SetIsReversed(v bool) *ListWorkspacesRequest {
	s.IsReversed = &v
	return s
}

func (s *ListWorkspacesRequest) SetLabelSelector(v string) *ListWorkspacesRequest {
	s.LabelSelector = &v
	return s
}

func (s *ListWorkspacesRequest) SetMaxResults(v int32) *ListWorkspacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesRequest) SetNextToken(v string) *ListWorkspacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesRequest) SetOrderBy(v string) *ListWorkspacesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListWorkspacesRequest) SetSearch(v string) *ListWorkspacesRequest {
	s.Search = &v
	return s
}

type ListWorkspacesResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// TestToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Workspaces []*ListWorkspacesResponseBodyWorkspaces `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) SetHostId(v string) *ListWorkspacesResponseBody {
	s.HostId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetMaxResults(v int32) *ListWorkspacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetNextToken(v string) *ListWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalCount(v int32) *ListWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

type ListWorkspacesResponseBodyWorkspaces struct {
	// example:
	//
	// my-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// This is a test workspace
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 30
	JobLifecycle *int32             `json:"JobLifecycle,omitempty" xml:"JobLifecycle,omitempty"`
	Labels       map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// 2021-06-16T15:09:14.378297+08:00
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// example:
	//
	// cn-beijing
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// RAM Role
	//
	// example:
	//
	// TestRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// oss://my-bucket/my-path/
	Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaces) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetBucketName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.BucketName = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreateTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDescription(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Description = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetJobLifecycle(v int32) *ListWorkspacesResponseBodyWorkspaces {
	s.JobLifecycle = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetLabels(v map[string]*string) *ListWorkspacesResponseBodyWorkspaces {
	s.Labels = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetLastModifiedTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.LastModifiedTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetLocation(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Location = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetRole(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Role = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetStatus(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Status = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetStorage(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Storage = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspace(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Workspace = &v
	return s
}

type ListWorkspacesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponse) SetHeaders(v map[string]*string) *ListWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspacesResponse) SetStatusCode(v int32) *ListWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspacesResponse) SetBody(v *ListWorkspacesResponseBody) *ListWorkspacesResponse {
	s.Body = v
	return s
}

type ResumeSubmissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sub-13BeXWnGvrOl3MC124tazcJVKaV
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ResumeSubmissionRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeSubmissionRequest) GoString() string {
	return s.String()
}

func (s *ResumeSubmissionRequest) SetSubmissionId(v string) *ResumeSubmissionRequest {
	s.SubmissionId = &v
	return s
}

func (s *ResumeSubmissionRequest) SetWorkspace(v string) *ResumeSubmissionRequest {
	s.Workspace = &v
	return s
}

type ResumeSubmissionResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeSubmissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeSubmissionResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeSubmissionResponseBody) SetHostId(v string) *ResumeSubmissionResponseBody {
	s.HostId = &v
	return s
}

func (s *ResumeSubmissionResponseBody) SetRequestId(v string) *ResumeSubmissionResponseBody {
	s.RequestId = &v
	return s
}

type ResumeSubmissionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeSubmissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeSubmissionResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeSubmissionResponse) GoString() string {
	return s.String()
}

func (s *ResumeSubmissionResponse) SetHeaders(v map[string]*string) *ResumeSubmissionResponse {
	s.Headers = v
	return s
}

func (s *ResumeSubmissionResponse) SetStatusCode(v int32) *ResumeSubmissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeSubmissionResponse) SetBody(v *ResumeSubmissionResponseBody) *ResumeSubmissionResponse {
	s.Body = v
	return s
}

type TagAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppRevision *string `json:"AppRevision,omitempty" xml:"AppRevision,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s TagAppRequest) String() string {
	return tea.Prettify(s)
}

func (s TagAppRequest) GoString() string {
	return s.String()
}

func (s *TagAppRequest) SetAppName(v string) *TagAppRequest {
	s.AppName = &v
	return s
}

func (s *TagAppRequest) SetAppRevision(v string) *TagAppRequest {
	s.AppRevision = &v
	return s
}

func (s *TagAppRequest) SetRevisionTag(v string) *TagAppRequest {
	s.RevisionTag = &v
	return s
}

func (s *TagAppRequest) SetWorkspace(v string) *TagAppRequest {
	s.Workspace = &v
	return s
}

type TagAppResponseBody struct {
	// example:
	//
	// TestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	AppRevision *string `json:"AppRevision,omitempty" xml:"AppRevision,omitempty"`
	// example:
	//
	// 0
	FormerRevision *string `json:"FormerRevision,omitempty" xml:"FormerRevision,omitempty"`
	// example:
	//
	// TestTag
	FormerTag *string `json:"FormerTag,omitempty" xml:"FormerTag,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ProduceVersion
	RevisionTag *string `json:"RevisionTag,omitempty" xml:"RevisionTag,omitempty"`
	// example:
	//
	// TestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s TagAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagAppResponseBody) GoString() string {
	return s.String()
}

func (s *TagAppResponseBody) SetAppName(v string) *TagAppResponseBody {
	s.AppName = &v
	return s
}

func (s *TagAppResponseBody) SetAppRevision(v string) *TagAppResponseBody {
	s.AppRevision = &v
	return s
}

func (s *TagAppResponseBody) SetFormerRevision(v string) *TagAppResponseBody {
	s.FormerRevision = &v
	return s
}

func (s *TagAppResponseBody) SetFormerTag(v string) *TagAppResponseBody {
	s.FormerTag = &v
	return s
}

func (s *TagAppResponseBody) SetRequestId(v string) *TagAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagAppResponseBody) SetRevisionTag(v string) *TagAppResponseBody {
	s.RevisionTag = &v
	return s
}

func (s *TagAppResponseBody) SetWorkspace(v string) *TagAppResponseBody {
	s.Workspace = &v
	return s
}

type TagAppResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagAppResponse) String() string {
	return tea.Prettify(s)
}

func (s TagAppResponse) GoString() string {
	return s.String()
}

func (s *TagAppResponse) SetHeaders(v map[string]*string) *TagAppResponse {
	s.Headers = v
	return s
}

func (s *TagAppResponse) SetStatusCode(v int32) *TagAppResponse {
	s.StatusCode = &v
	return s
}

func (s *TagAppResponse) SetBody(v *TagAppResponseBody) *TagAppResponse {
	s.Body = v
	return s
}

type UpdateEntityRequest struct {
	// This parameter is required.
	EntityItems []*UpdateEntityRequestEntityItems `json:"EntityItems,omitempty" xml:"EntityItems,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UpdateEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityRequest) GoString() string {
	return s.String()
}

func (s *UpdateEntityRequest) SetEntityItems(v []*UpdateEntityRequestEntityItems) *UpdateEntityRequest {
	s.EntityItems = v
	return s
}

func (s *UpdateEntityRequest) SetEntityType(v string) *UpdateEntityRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateEntityRequest) SetWorkspace(v string) *UpdateEntityRequest {
	s.Workspace = &v
	return s
}

type UpdateEntityRequestEntityItems struct {
	EntityData map[string]*string `json:"EntityData,omitempty" xml:"EntityData,omitempty"`
	// example:
	//
	// s1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
}

func (s UpdateEntityRequestEntityItems) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityRequestEntityItems) GoString() string {
	return s.String()
}

func (s *UpdateEntityRequestEntityItems) SetEntityData(v map[string]*string) *UpdateEntityRequestEntityItems {
	s.EntityData = v
	return s
}

func (s *UpdateEntityRequestEntityItems) SetEntityName(v string) *UpdateEntityRequestEntityItems {
	s.EntityName = &v
	return s
}

type UpdateEntityShrinkRequest struct {
	// This parameter is required.
	EntityItemsShrink *string `json:"EntityItems,omitempty" xml:"EntityItems,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UpdateEntityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEntityShrinkRequest) SetEntityItemsShrink(v string) *UpdateEntityShrinkRequest {
	s.EntityItemsShrink = &v
	return s
}

func (s *UpdateEntityShrinkRequest) SetEntityType(v string) *UpdateEntityShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateEntityShrinkRequest) SetWorkspace(v string) *UpdateEntityShrinkRequest {
	s.Workspace = &v
	return s
}

type UpdateEntityResponseBody struct {
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UpdateEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEntityResponseBody) SetEntityType(v string) *UpdateEntityResponseBody {
	s.EntityType = &v
	return s
}

func (s *UpdateEntityResponseBody) SetHostId(v string) *UpdateEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *UpdateEntityResponseBody) SetRequestId(v string) *UpdateEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEntityResponseBody) SetWorkspace(v string) *UpdateEntityResponseBody {
	s.Workspace = &v
	return s
}

type UpdateEntityResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityResponse) GoString() string {
	return s.String()
}

func (s *UpdateEntityResponse) SetHeaders(v map[string]*string) *UpdateEntityResponse {
	s.Headers = v
	return s
}

func (s *UpdateEntityResponse) SetStatusCode(v int32) *UpdateEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEntityResponse) SetBody(v *UpdateEntityResponseBody) *UpdateEntityResponse {
	s.Body = v
	return s
}

type UpdateEntityItemsRequest struct {
	// This parameter is required.
	EntityItems []*UpdateEntityItemsRequestEntityItems `json:"EntityItems,omitempty" xml:"EntityItems,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UpdateEntityItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityItemsRequest) GoString() string {
	return s.String()
}

func (s *UpdateEntityItemsRequest) SetEntityItems(v []*UpdateEntityItemsRequestEntityItems) *UpdateEntityItemsRequest {
	s.EntityItems = v
	return s
}

func (s *UpdateEntityItemsRequest) SetEntityType(v string) *UpdateEntityItemsRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateEntityItemsRequest) SetWorkspace(v string) *UpdateEntityItemsRequest {
	s.Workspace = &v
	return s
}

type UpdateEntityItemsRequestEntityItems struct {
	EntityData map[string]*string `json:"EntityData,omitempty" xml:"EntityData,omitempty"`
	// example:
	//
	// s1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
}

func (s UpdateEntityItemsRequestEntityItems) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityItemsRequestEntityItems) GoString() string {
	return s.String()
}

func (s *UpdateEntityItemsRequestEntityItems) SetEntityData(v map[string]*string) *UpdateEntityItemsRequestEntityItems {
	s.EntityData = v
	return s
}

func (s *UpdateEntityItemsRequestEntityItems) SetEntityName(v string) *UpdateEntityItemsRequestEntityItems {
	s.EntityName = &v
	return s
}

type UpdateEntityItemsShrinkRequest struct {
	// This parameter is required.
	EntityItemsShrink *string `json:"EntityItems,omitempty" xml:"EntityItems,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UpdateEntityItemsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEntityItemsShrinkRequest) SetEntityItemsShrink(v string) *UpdateEntityItemsShrinkRequest {
	s.EntityItemsShrink = &v
	return s
}

func (s *UpdateEntityItemsShrinkRequest) SetEntityType(v string) *UpdateEntityItemsShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateEntityItemsShrinkRequest) SetWorkspace(v string) *UpdateEntityItemsShrinkRequest {
	s.Workspace = &v
	return s
}

type UpdateEntityItemsResponseBody struct {
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UpdateEntityItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityItemsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEntityItemsResponseBody) SetEntityType(v string) *UpdateEntityItemsResponseBody {
	s.EntityType = &v
	return s
}

func (s *UpdateEntityItemsResponseBody) SetHostId(v string) *UpdateEntityItemsResponseBody {
	s.HostId = &v
	return s
}

func (s *UpdateEntityItemsResponseBody) SetRequestId(v string) *UpdateEntityItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEntityItemsResponseBody) SetWorkspace(v string) *UpdateEntityItemsResponseBody {
	s.Workspace = &v
	return s
}

type UpdateEntityItemsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEntityItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEntityItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityItemsResponse) GoString() string {
	return s.String()
}

func (s *UpdateEntityItemsResponse) SetHeaders(v map[string]*string) *UpdateEntityItemsResponse {
	s.Headers = v
	return s
}

func (s *UpdateEntityItemsResponse) SetStatusCode(v int32) *UpdateEntityItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEntityItemsResponse) SetBody(v *UpdateEntityItemsResponseBody) *UpdateEntityItemsResponse {
	s.Body = v
	return s
}

type UpdateTemplateRequest struct {
	// example:
	//
	// This is a test template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InputsExpression []*UpdateTemplateRequestInputsExpression `json:"InputsExpression,omitempty" xml:"InputsExpression,omitempty" type:"Repeated"`
	// example:
	//
	// {"key": "value"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// This parameter is required.
	OutputsExpression []*UpdateTemplateRequestOutputsExpression `json:"OutputsExpression,omitempty" xml:"OutputsExpression,omitempty" type:"Repeated"`
	// example:
	//
	// test-entity
	RootEntity *string `json:"RootEntity,omitempty" xml:"RootEntity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) SetDescription(v string) *UpdateTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateRequest) SetInputsExpression(v []*UpdateTemplateRequestInputsExpression) *UpdateTemplateRequest {
	s.InputsExpression = v
	return s
}

func (s *UpdateTemplateRequest) SetLabels(v string) *UpdateTemplateRequest {
	s.Labels = &v
	return s
}

func (s *UpdateTemplateRequest) SetOutputsExpression(v []*UpdateTemplateRequestOutputsExpression) *UpdateTemplateRequest {
	s.OutputsExpression = v
	return s
}

func (s *UpdateTemplateRequest) SetRootEntity(v string) *UpdateTemplateRequest {
	s.RootEntity = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateName(v string) *UpdateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateRequest) SetWorkspace(v string) *UpdateTemplateRequest {
	s.Workspace = &v
	return s
}

type UpdateTemplateRequestInputsExpression struct {
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// True
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// 2
	StepOrder *int32 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fastq1
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// String
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// read_1.fastq.gz
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s UpdateTemplateRequestInputsExpression) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequestInputsExpression) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequestInputsExpression) SetHelp(v string) *UpdateTemplateRequestInputsExpression {
	s.Help = &v
	return s
}

func (s *UpdateTemplateRequestInputsExpression) SetRequired(v bool) *UpdateTemplateRequestInputsExpression {
	s.Required = &v
	return s
}

func (s *UpdateTemplateRequestInputsExpression) SetStepOrder(v int32) *UpdateTemplateRequestInputsExpression {
	s.StepOrder = &v
	return s
}

func (s *UpdateTemplateRequestInputsExpression) SetTaskName(v string) *UpdateTemplateRequestInputsExpression {
	s.TaskName = &v
	return s
}

func (s *UpdateTemplateRequestInputsExpression) SetVariableName(v string) *UpdateTemplateRequestInputsExpression {
	s.VariableName = &v
	return s
}

func (s *UpdateTemplateRequestInputsExpression) SetVariableType(v string) *UpdateTemplateRequestInputsExpression {
	s.VariableType = &v
	return s
}

func (s *UpdateTemplateRequestInputsExpression) SetVariableValue(v string) *UpdateTemplateRequestInputsExpression {
	s.VariableValue = &v
	return s
}

type UpdateTemplateRequestOutputsExpression struct {
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// True
	Required  *bool  `json:"Required,omitempty" xml:"Required,omitempty"`
	StepOrder *int32 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bam
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// String
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// out.bam
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s UpdateTemplateRequestOutputsExpression) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequestOutputsExpression) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequestOutputsExpression) SetHelp(v string) *UpdateTemplateRequestOutputsExpression {
	s.Help = &v
	return s
}

func (s *UpdateTemplateRequestOutputsExpression) SetRequired(v bool) *UpdateTemplateRequestOutputsExpression {
	s.Required = &v
	return s
}

func (s *UpdateTemplateRequestOutputsExpression) SetStepOrder(v int32) *UpdateTemplateRequestOutputsExpression {
	s.StepOrder = &v
	return s
}

func (s *UpdateTemplateRequestOutputsExpression) SetTaskName(v string) *UpdateTemplateRequestOutputsExpression {
	s.TaskName = &v
	return s
}

func (s *UpdateTemplateRequestOutputsExpression) SetVariableName(v string) *UpdateTemplateRequestOutputsExpression {
	s.VariableName = &v
	return s
}

func (s *UpdateTemplateRequestOutputsExpression) SetVariableType(v string) *UpdateTemplateRequestOutputsExpression {
	s.VariableType = &v
	return s
}

func (s *UpdateTemplateRequestOutputsExpression) SetVariableValue(v string) *UpdateTemplateRequestOutputsExpression {
	s.VariableValue = &v
	return s
}

type UpdateTemplateShrinkRequest struct {
	// example:
	//
	// This is a test template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InputsExpressionShrink *string `json:"InputsExpression,omitempty" xml:"InputsExpression,omitempty"`
	// example:
	//
	// {"key": "value"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// This parameter is required.
	OutputsExpressionShrink *string `json:"OutputsExpression,omitempty" xml:"OutputsExpression,omitempty"`
	// example:
	//
	// test-entity
	RootEntity *string `json:"RootEntity,omitempty" xml:"RootEntity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UpdateTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateShrinkRequest) SetDescription(v string) *UpdateTemplateShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetInputsExpressionShrink(v string) *UpdateTemplateShrinkRequest {
	s.InputsExpressionShrink = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetLabels(v string) *UpdateTemplateShrinkRequest {
	s.Labels = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetOutputsExpressionShrink(v string) *UpdateTemplateShrinkRequest {
	s.OutputsExpressionShrink = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetRootEntity(v string) *UpdateTemplateShrinkRequest {
	s.RootEntity = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetTemplateName(v string) *UpdateTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetWorkspace(v string) *UpdateTemplateShrinkRequest {
	s.Workspace = &v
	return s
}

type UpdateTemplateResponseBody struct {
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) SetHostId(v string) *UpdateTemplateResponseBody {
	s.HostId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateResponse) SetStatusCode(v int32) *UpdateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTemplateResponse) SetBody(v *UpdateTemplateResponseBody) *UpdateTemplateResponse {
	s.Body = v
	return s
}

type UpdateWorkspaceRequest struct {
	// example:
	//
	// This is a test workspace
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 30
	JobLifecycle *int32 `json:"JobLifecycle,omitempty" xml:"JobLifecycle,omitempty"`
	// example:
	//
	// {"env":"test"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// MyTestRamRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MyTestWorkspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UpdateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequest) SetDescription(v string) *UpdateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetJobLifecycle(v int32) *UpdateWorkspaceRequest {
	s.JobLifecycle = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetLabels(v string) *UpdateWorkspaceRequest {
	s.Labels = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetRole(v string) *UpdateWorkspaceRequest {
	s.Role = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetWorkspace(v string) *UpdateWorkspaceRequest {
	s.Workspace = &v
	return s
}

type UpdateWorkspaceResponseBody struct {
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBody) SetHostId(v string) *UpdateWorkspaceResponseBody {
	s.HostId = &v
	return s
}

func (s *UpdateWorkspaceResponseBody) SetRequestId(v string) *UpdateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceResponse) SetStatusCode(v int32) *UpdateWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceResponse) SetBody(v *UpdateWorkspaceResponseBody) *UpdateWorkspaceResponse {
	s.Body = v
	return s
}

type UploadEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://my-bucket.oss.aliyucs.com/entity.csv
	EntityCSVFile *string `json:"EntityCSVFile,omitempty" xml:"EntityCSVFile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UploadEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadEntityRequest) GoString() string {
	return s.String()
}

func (s *UploadEntityRequest) SetEntityCSVFile(v string) *UploadEntityRequest {
	s.EntityCSVFile = &v
	return s
}

func (s *UploadEntityRequest) SetWorkspace(v string) *UploadEntityRequest {
	s.Workspace = &v
	return s
}

type UploadEntityResponseBody struct {
	// example:
	//
	// sample
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// easygene.cn-beijing.aliyuncs.com
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// DA980AD0-158F-44F3-847D-5EAB96C0EB6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test-workspace
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UploadEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UploadEntityResponseBody) SetEntityType(v string) *UploadEntityResponseBody {
	s.EntityType = &v
	return s
}

func (s *UploadEntityResponseBody) SetHostId(v string) *UploadEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *UploadEntityResponseBody) SetRequestId(v string) *UploadEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadEntityResponseBody) SetWorkspace(v string) *UploadEntityResponseBody {
	s.Workspace = &v
	return s
}

type UploadEntityResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadEntityResponse) GoString() string {
	return s.String()
}

func (s *UploadEntityResponse) SetHeaders(v map[string]*string) *UploadEntityResponse {
	s.Headers = v
	return s
}

func (s *UploadEntityResponse) SetStatusCode(v int32) *UploadEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadEntityResponse) SetBody(v *UploadEntityResponseBody) *UploadEntityResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("easygene"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 暂停任务
//
// @param request - AbortRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbortRunResponse
func (client *Client) AbortRunWithOptions(request *AbortRunRequest, runtime *util.RuntimeOptions) (_result *AbortRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RunId)) {
		query["RunId"] = request.RunId
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AbortRun"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AbortRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停任务
//
// @param request - AbortRunRequest
//
// @return AbortRunResponse
func (client *Client) AbortRun(request *AbortRunRequest) (_result *AbortRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbortRunResponse{}
	_body, _err := client.AbortRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消投递
//
// @param request - AbortSubmissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbortSubmissionResponse
func (client *Client) AbortSubmissionWithOptions(request *AbortSubmissionRequest, runtime *util.RuntimeOptions) (_result *AbortSubmissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubmissionId)) {
		query["SubmissionId"] = request.SubmissionId
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AbortSubmission"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AbortSubmissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消投递
//
// @param request - AbortSubmissionRequest
//
// @return AbortSubmissionResponse
func (client *Client) AbortSubmission(request *AbortSubmissionRequest) (_result *AbortSubmissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbortSubmissionResponse{}
	_body, _err := client.AbortSubmissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拷贝公共数据集的数据表格
//
// @param request - CopyPublicEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyPublicEntityResponse
func (client *Client) CopyPublicEntityWithOptions(request *CopyPublicEntityRequest, runtime *util.RuntimeOptions) (_result *CopyPublicEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dataset)) {
		query["Dataset"] = request.Dataset
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyPublicEntity"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyPublicEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拷贝公共数据集的数据表格
//
// @param request - CopyPublicEntityRequest
//
// @return CopyPublicEntityResponse
func (client *Client) CopyPublicEntity(request *CopyPublicEntityRequest) (_result *CopyPublicEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyPublicEntityResponse{}
	_body, _err := client.CopyPublicEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param tmpReq - CreateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithOptions(tmpReq *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Configs)) {
		request.ConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configs, tea.String("Configs"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Dependencies)) {
		request.DependenciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dependencies, tea.String("Dependencies"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageVersion)) {
		query["LanguageVersion"] = request.LanguageVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionComment)) {
		query["RevisionComment"] = request.RevisionComment
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionTag)) {
		query["RevisionTag"] = request.RevisionTag
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigsShrink)) {
		body["Configs"] = request.ConfigsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Definition)) {
		body["Definition"] = request.Definition
	}

	if !tea.BoolValue(util.IsUnset(request.DependenciesShrink)) {
		body["Dependencies"] = request.DependenciesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Documentation)) {
		body["Documentation"] = request.Documentation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param request - CreateAppRequest
//
// @return CreateAppResponse
func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据表格
//
// @param tmpReq - CreateEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEntityResponse
func (client *Client) CreateEntityWithOptions(tmpReq *CreateEntityRequest, runtime *util.RuntimeOptions) (_result *CreateEntityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityItems)) {
		request.EntityItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityItems, tea.String("EntityItems"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityItemsShrink)) {
		body["EntityItems"] = request.EntityItemsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEntity"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据表格
//
// @param request - CreateEntityRequest
//
// @return CreateEntityResponse
func (client *Client) CreateEntity(request *CreateEntityRequest) (_result *CreateEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEntityResponse{}
	_body, _err := client.CreateEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建任务
//
// @param tmpReq - CreateRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRunResponse
func (client *Client) CreateRunWithOptions(tmpReq *CreateRunRequest, runtime *util.RuntimeOptions) (_result *CreateRunResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateRunShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExecuteOptions)) {
		request.ExecuteOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecuteOptions, tea.String("ExecuteOptions"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppRevision)) {
		query["AppRevision"] = request.AppRevision
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRuntime)) {
		query["DefaultRuntime"] = request.DefaultRuntime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteDirectory)) {
		query["ExecuteDirectory"] = request.ExecuteDirectory
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteOptionsShrink)) {
		query["ExecuteOptions"] = request.ExecuteOptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Inputs)) {
		query["Inputs"] = request.Inputs
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFolder)) {
		query["OutputFolder"] = request.OutputFolder
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionTag)) {
		query["RevisionTag"] = request.RevisionTag
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.RunName)) {
		query["RunName"] = request.RunName
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRun"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建任务
//
// @param request - CreateRunRequest
//
// @return CreateRunResponse
func (client *Client) CreateRun(request *CreateRunRequest) (_result *CreateRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRunResponse{}
	_body, _err := client.CreateRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建提交
//
// @param tmpReq - CreateSubmissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubmissionResponse
func (client *Client) CreateSubmissionWithOptions(tmpReq *CreateSubmissionRequest, runtime *util.RuntimeOptions) (_result *CreateSubmissionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSubmissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityNames)) {
		request.EntityNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityNames, tea.String("EntityNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRuntime)) {
		query["DefaultRuntime"] = request.DefaultRuntime
	}

	if !tea.BoolValue(util.IsUnset(request.EntityNamesShrink)) {
		query["EntityNames"] = request.EntityNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteDirectory)) {
		query["ExecuteDirectory"] = request.ExecuteDirectory
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteOptions)) {
		query["ExecuteOptions"] = request.ExecuteOptions
	}

	if !tea.BoolValue(util.IsUnset(request.Inputs)) {
		query["Inputs"] = request.Inputs
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFolder)) {
		query["OutputFolder"] = request.OutputFolder
	}

	if !tea.BoolValue(util.IsUnset(request.Outputs)) {
		query["Outputs"] = request.Outputs
	}

	if !tea.BoolValue(util.IsUnset(request.Revision)) {
		query["Revision"] = request.Revision
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionTag)) {
		query["RevisionTag"] = request.RevisionTag
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSubmission"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSubmissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建提交
//
// @param request - CreateSubmissionRequest
//
// @return CreateSubmissionResponse
func (client *Client) CreateSubmission(request *CreateSubmissionRequest) (_result *CreateSubmissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSubmissionResponse{}
	_body, _err := client.CreateSubmissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建应用模板
//
// @param tmpReq - CreateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithOptions(tmpReq *CreateTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InputsExpression)) {
		request.InputsExpressionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputsExpression, tea.String("InputsExpression"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OutputsExpression)) {
		request.OutputsExpressionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutputsExpression, tea.String("OutputsExpression"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppRevision)) {
		query["AppRevision"] = request.AppRevision
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionTag)) {
		query["RevisionTag"] = request.RevisionTag
	}

	if !tea.BoolValue(util.IsUnset(request.RootEntity)) {
		query["RootEntity"] = request.RootEntity
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InputsExpressionShrink)) {
		body["InputsExpression"] = request.InputsExpressionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OutputsExpressionShrink)) {
		body["OutputsExpression"] = request.OutputsExpressionShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTemplate"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用模板
//
// @param request - CreateTemplateRequest
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建工作空间
//
// @param request - CreateWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspaceWithOptions(request *CreateWorkspaceRequest, runtime *util.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.JobLifecycle)) {
		query["JobLifecycle"] = request.JobLifecycle
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Storage)) {
		query["Storage"] = request.Storage
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkspace"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工作空间
//
// @param request - CreateWorkspaceRequest
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除应用
//
// @param request - DeleteAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Revision)) {
		query["Revision"] = request.Revision
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApp"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用
//
// @param request - DeleteAppRequest
//
// @return DeleteAppResponse
func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除实体表格
//
// @param request - DeleteEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEntityResponse
func (client *Client) DeleteEntityWithOptions(request *DeleteEntityRequest, runtime *util.RuntimeOptions) (_result *DeleteEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEntity"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实体表格
//
// @param request - DeleteEntityRequest
//
// @return DeleteEntityResponse
func (client *Client) DeleteEntity(request *DeleteEntityRequest) (_result *DeleteEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEntityResponse{}
	_body, _err := client.DeleteEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据表格的条目
//
// @param tmpReq - DeleteEntityItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEntityItemsResponse
func (client *Client) DeleteEntityItemsWithOptions(tmpReq *DeleteEntityItemsRequest, runtime *util.RuntimeOptions) (_result *DeleteEntityItemsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteEntityItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityNames)) {
		request.EntityNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityNames, tea.String("EntityNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityNamesShrink)) {
		query["EntityNames"] = request.EntityNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEntityItems"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEntityItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据表格的条目
//
// @param request - DeleteEntityItemsRequest
//
// @return DeleteEntityItemsResponse
func (client *Client) DeleteEntityItems(request *DeleteEntityItemsRequest) (_result *DeleteEntityItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEntityItemsResponse{}
	_body, _err := client.DeleteEntityItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除任务
//
// @param request - DeleteRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRunResponse
func (client *Client) DeleteRunWithOptions(request *DeleteRunRequest, runtime *util.RuntimeOptions) (_result *DeleteRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RunId)) {
		query["RunId"] = request.RunId
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRun"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除任务
//
// @param request - DeleteRunRequest
//
// @return DeleteRunResponse
func (client *Client) DeleteRun(request *DeleteRunRequest) (_result *DeleteRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRunResponse{}
	_body, _err := client.DeleteRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除投递
//
// @param request - DeleteSubmissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubmissionResponse
func (client *Client) DeleteSubmissionWithOptions(request *DeleteSubmissionRequest, runtime *util.RuntimeOptions) (_result *DeleteSubmissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubmissionId)) {
		query["SubmissionId"] = request.SubmissionId
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSubmission"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSubmissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除投递
//
// @param request - DeleteSubmissionRequest
//
// @return DeleteSubmissionResponse
func (client *Client) DeleteSubmission(request *DeleteSubmissionRequest) (_result *DeleteSubmissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSubmissionResponse{}
	_body, _err := client.DeleteSubmissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除应用模板
//
// @param request - DeleteTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTemplate"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用模板
//
// @param request - DeleteTemplateRequest
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除工作空间
//
// @param request - DeleteWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithOptions(request *DeleteWorkspaceRequest, runtime *util.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkspace"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除工作空间
//
// @param request - DeleteWorkspaceRequest
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspace(request *DeleteWorkspaceRequest) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下载数据表格
//
// @param tmpReq - DownloadEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadEntityResponse
func (client *Client) DownloadEntityWithOptions(tmpReq *DownloadEntityRequest, runtime *util.RuntimeOptions) (_result *DownloadEntityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DownloadEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityNames)) {
		request.EntityNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityNames, tea.String("EntityNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityNamesShrink)) {
		query["EntityNames"] = request.EntityNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadEntity"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载数据表格
//
// @param request - DownloadEntityRequest
//
// @return DownloadEntityResponse
func (client *Client) DownloadEntity(request *DownloadEntityRequest) (_result *DownloadEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadEntityResponse{}
	_body, _err := client.DownloadEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询应用详情
//
// @param request - GetAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppResponse
func (client *Client) GetAppWithOptions(request *GetAppRequest, runtime *util.RuntimeOptions) (_result *GetAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Revision)) {
		query["Revision"] = request.Revision
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionTag)) {
		query["RevisionTag"] = request.RevisionTag
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApp"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用详情
//
// @param request - GetAppRequest
//
// @return GetAppResponse
func (client *Client) GetApp(request *GetAppRequest) (_result *GetAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppResponse{}
	_body, _err := client.GetAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据表格信息
//
// @param request - GetEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEntityResponse
func (client *Client) GetEntityWithOptions(request *GetEntityRequest, runtime *util.RuntimeOptions) (_result *GetEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEntity"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据表格信息
//
// @param request - GetEntityRequest
//
// @return GetEntityResponse
func (client *Client) GetEntity(request *GetEntityRequest) (_result *GetEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEntityResponse{}
	_body, _err := client.GetEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询平台公共应用详情
//
// @param tmpReq - GetGlobalAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGlobalAppResponse
func (client *Client) GetGlobalAppWithOptions(tmpReq *GetGlobalAppRequest, runtime *util.RuntimeOptions) (_result *GetGlobalAppResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetGlobalAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Attributes)) {
		request.AttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attributes, tea.String("Attributes"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGlobalApp"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGlobalAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询平台公共应用详情
//
// @param request - GetGlobalAppRequest
//
// @return GetGlobalAppResponse
func (client *Client) GetGlobalApp(request *GetGlobalAppRequest) (_result *GetGlobalAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGlobalAppResponse{}
	_body, _err := client.GetGlobalAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询平台公共数据集详情
//
// @param tmpReq - GetPublicDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublicDatasetResponse
func (client *Client) GetPublicDatasetWithOptions(tmpReq *GetPublicDatasetRequest, runtime *util.RuntimeOptions) (_result *GetPublicDatasetResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetPublicDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Attributes)) {
		request.AttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attributes, tea.String("Attributes"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPublicDataset"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPublicDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询平台公共数据集详情
//
// @param request - GetPublicDatasetRequest
//
// @return GetPublicDatasetResponse
func (client *Client) GetPublicDataset(request *GetPublicDatasetRequest) (_result *GetPublicDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPublicDatasetResponse{}
	_body, _err := client.GetPublicDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询平台公共数据集特定的实体定义
//
// @param request - GetPublicDatasetEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublicDatasetEntityResponse
func (client *Client) GetPublicDatasetEntityWithOptions(request *GetPublicDatasetEntityRequest, runtime *util.RuntimeOptions) (_result *GetPublicDatasetEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPublicDatasetEntity"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPublicDatasetEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询平台公共数据集特定的实体定义
//
// @param request - GetPublicDatasetEntityRequest
//
// @return GetPublicDatasetEntityResponse
func (client *Client) GetPublicDatasetEntity(request *GetPublicDatasetEntityRequest) (_result *GetPublicDatasetEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPublicDatasetEntityResponse{}
	_body, _err := client.GetPublicDatasetEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务详情
//
// @param request - GetRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRunResponse
func (client *Client) GetRunWithOptions(request *GetRunRequest, runtime *util.RuntimeOptions) (_result *GetRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRun"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务详情
//
// @param request - GetRunRequest
//
// @return GetRunResponse
func (client *Client) GetRun(request *GetRunRequest) (_result *GetRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRunResponse{}
	_body, _err := client.GetRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取投递详情
//
// @param request - GetSubmissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubmissionResponse
func (client *Client) GetSubmissionWithOptions(request *GetSubmissionRequest, runtime *util.RuntimeOptions) (_result *GetSubmissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubmission"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubmissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取投递详情
//
// @param request - GetSubmissionRequest
//
// @return GetSubmissionResponse
func (client *Client) GetSubmission(request *GetSubmissionRequest) (_result *GetSubmissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubmissionResponse{}
	_body, _err := client.GetSubmissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询应用模板的详情
//
// @param request - GetTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplate"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用模板的详情
//
// @param request - GetTemplateRequest
//
// @return GetTemplateResponse
func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询工作空间详情
//
// @param request - GetWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithOptions(request *GetWorkspaceRequest, runtime *util.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkspace"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工作空间详情
//
// @param request - GetWorkspaceRequest
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspace(request *GetWorkspaceRequest) (_result *GetWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入应用。
//
// @param request - ImportAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportAppResponse
func (client *Client) ImportAppWithOptions(request *ImportAppRequest, runtime *util.RuntimeOptions) (_result *ImportAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportApp"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入应用。
//
// @param request - ImportAppRequest
//
// @return ImportAppResponse
func (client *Client) ImportApp(request *ImportAppRequest) (_result *ImportAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportAppResponse{}
	_body, _err := client.ImportAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 安装平台公共应用到工作空间中。
//
// @param request - InstallGlobalAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallGlobalAppResponse
func (client *Client) InstallGlobalAppWithOptions(request *InstallGlobalAppRequest, runtime *util.RuntimeOptions) (_result *InstallGlobalAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.InstalledAppName)) {
		query["InstalledAppName"] = request.InstalledAppName
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallGlobalApp"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallGlobalAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 安装平台公共应用到工作空间中。
//
// @param request - InstallGlobalAppRequest
//
// @return InstallGlobalAppResponse
func (client *Client) InstallGlobalApp(request *InstallGlobalAppRequest) (_result *InstallGlobalAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallGlobalAppResponse{}
	_body, _err := client.InstallGlobalAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出应用
//
// @param request - ListAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppsResponse
func (client *Client) ListAppsWithOptions(request *ListAppsRequest, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.IsReversed)) {
		query["IsReversed"] = request.IsReversed
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelector)) {
		query["LabelSelector"] = request.LabelSelector
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApps"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出应用
//
// @param request - ListAppsRequest
//
// @return ListAppsResponse
func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取平台第三方软件列表
//
// @param request - ListAuthorizedSoftwareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedSoftwareResponse
func (client *Client) ListAuthorizedSoftwareWithOptions(request *ListAuthorizedSoftwareRequest, runtime *util.RuntimeOptions) (_result *ListAuthorizedSoftwareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAuthorizedSoftware"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAuthorizedSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取平台第三方软件列表
//
// @param request - ListAuthorizedSoftwareRequest
//
// @return ListAuthorizedSoftwareResponse
func (client *Client) ListAuthorizedSoftware(request *ListAuthorizedSoftwareRequest) (_result *ListAuthorizedSoftwareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAuthorizedSoftwareResponse{}
	_body, _err := client.ListAuthorizedSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取平台公共容器镜像列表
//
// @param request - ListContainerImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContainerImagesResponse
func (client *Client) ListContainerImagesWithOptions(request *ListContainerImagesRequest, runtime *util.RuntimeOptions) (_result *ListContainerImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListContainerImages"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListContainerImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取平台公共容器镜像列表
//
// @param request - ListContainerImagesRequest
//
// @return ListContainerImagesResponse
func (client *Client) ListContainerImages(request *ListContainerImagesRequest) (_result *ListContainerImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContainerImagesResponse{}
	_body, _err := client.ListContainerImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出数据实体列表
//
// @param request - ListEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEntitiesResponse
func (client *Client) ListEntitiesWithOptions(request *ListEntitiesRequest, runtime *util.RuntimeOptions) (_result *ListEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEntities"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出数据实体列表
//
// @param request - ListEntitiesRequest
//
// @return ListEntitiesResponse
func (client *Client) ListEntities(request *ListEntitiesRequest) (_result *ListEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEntitiesResponse{}
	_body, _err := client.ListEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出数据表格的条目
//
// @param request - ListEntityItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEntityItemsResponse
func (client *Client) ListEntityItemsWithOptions(request *ListEntityItemsRequest, runtime *util.RuntimeOptions) (_result *ListEntityItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.IsReversed)) {
		query["IsReversed"] = request.IsReversed
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEntityItems"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEntityItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出数据表格的条目
//
// @param request - ListEntityItemsRequest
//
// @return ListEntityItemsResponse
func (client *Client) ListEntityItems(request *ListEntityItemsRequest) (_result *ListEntityItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEntityItemsResponse{}
	_body, _err := client.ListEntityItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取平台公共应用列表
//
// @param request - ListGlobalAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGlobalAppsResponse
func (client *Client) ListGlobalAppsWithOptions(request *ListGlobalAppsRequest, runtime *util.RuntimeOptions) (_result *ListGlobalAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGlobalApps"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGlobalAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取平台公共应用列表
//
// @param request - ListGlobalAppsRequest
//
// @return ListGlobalAppsResponse
func (client *Client) ListGlobalApps(request *ListGlobalAppsRequest) (_result *ListGlobalAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGlobalAppsResponse{}
	_body, _err := client.ListGlobalAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取平台公共数据集列表
//
// @param request - ListPublicDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublicDatasetResponse
func (client *Client) ListPublicDatasetWithOptions(request *ListPublicDatasetRequest, runtime *util.RuntimeOptions) (_result *ListPublicDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublicDataset"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublicDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取平台公共数据集列表
//
// @param request - ListPublicDatasetRequest
//
// @return ListPublicDatasetResponse
func (client *Client) ListPublicDataset(request *ListPublicDatasetRequest) (_result *ListPublicDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicDatasetResponse{}
	_body, _err := client.ListPublicDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询平台公共数据集包含的实体列表
//
// @param request - ListPublicDatasetEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublicDatasetEntitiesResponse
func (client *Client) ListPublicDatasetEntitiesWithOptions(request *ListPublicDatasetEntitiesRequest, runtime *util.RuntimeOptions) (_result *ListPublicDatasetEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublicDatasetEntities"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublicDatasetEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询平台公共数据集包含的实体列表
//
// @param request - ListPublicDatasetEntitiesRequest
//
// @return ListPublicDatasetEntitiesResponse
func (client *Client) ListPublicDatasetEntities(request *ListPublicDatasetEntitiesRequest) (_result *ListPublicDatasetEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicDatasetEntitiesResponse{}
	_body, _err := client.ListPublicDatasetEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询平台公共数据集特定的实体表格数据
//
// @param request - ListPublicDatasetEntityItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublicDatasetEntityItemsResponse
func (client *Client) ListPublicDatasetEntityItemsWithOptions(request *ListPublicDatasetEntityItemsRequest, runtime *util.RuntimeOptions) (_result *ListPublicDatasetEntityItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublicDatasetEntityItems"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublicDatasetEntityItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询平台公共数据集特定的实体表格数据
//
// @param request - ListPublicDatasetEntityItemsRequest
//
// @return ListPublicDatasetEntityItemsResponse
func (client *Client) ListPublicDatasetEntityItems(request *ListPublicDatasetEntityItemsRequest) (_result *ListPublicDatasetEntityItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicDatasetEntityItemsResponse{}
	_body, _err := client.ListPublicDatasetEntityItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出所有公共数据集标签
//
// @param request - ListPublicDatasetTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublicDatasetTagsResponse
func (client *Client) ListPublicDatasetTagsWithOptions(request *ListPublicDatasetTagsRequest, runtime *util.RuntimeOptions) (_result *ListPublicDatasetTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublicDatasetTags"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublicDatasetTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出所有公共数据集标签
//
// @param request - ListPublicDatasetTagsRequest
//
// @return ListPublicDatasetTagsResponse
func (client *Client) ListPublicDatasetTags(request *ListPublicDatasetTagsRequest) (_result *ListPublicDatasetTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicDatasetTagsResponse{}
	_body, _err := client.ListPublicDatasetTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询基因分析平台产品的可用地域。
//
// @param request - ListRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询基因分析平台产品的可用地域。
//
// @return ListRegionsResponse
func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出任务
//
// @param request - ListRunsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRunsResponse
func (client *Client) ListRunsWithOptions(request *ListRunsRequest, runtime *util.RuntimeOptions) (_result *ListRunsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRuns"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRunsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出任务
//
// @param request - ListRunsRequest
//
// @return ListRunsResponse
func (client *Client) ListRuns(request *ListRunsRequest) (_result *ListRunsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRunsResponse{}
	_body, _err := client.ListRunsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出投递
//
// @param request - ListSubmissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubmissionsResponse
func (client *Client) ListSubmissionsWithOptions(request *ListSubmissionsRequest, runtime *util.RuntimeOptions) (_result *ListSubmissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubmissions"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubmissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出投递
//
// @param request - ListSubmissionsRequest
//
// @return ListSubmissionsResponse
func (client *Client) ListSubmissions(request *ListSubmissionsRequest) (_result *ListSubmissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSubmissionsResponse{}
	_body, _err := client.ListSubmissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出应用模板
//
// @param request - ListTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplatesResponse
func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsReversed)) {
		query["IsReversed"] = request.IsReversed
	}

	if !tea.BoolValue(util.IsUnset(request.LabelSelector)) {
		query["LabelSelector"] = request.LabelSelector
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplates"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出应用模板
//
// @param request - ListTemplatesRequest
//
// @return ListTemplatesResponse
func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 最近任务列表
//
// @param request - ListUserActiveRunsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserActiveRunsResponse
func (client *Client) ListUserActiveRunsWithOptions(request *ListUserActiveRunsRequest, runtime *util.RuntimeOptions) (_result *ListUserActiveRunsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserActiveRuns"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserActiveRunsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 最近任务列表
//
// @param request - ListUserActiveRunsRequest
//
// @return ListUserActiveRunsResponse
func (client *Client) ListUserActiveRuns(request *ListUserActiveRunsRequest) (_result *ListUserActiveRunsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserActiveRunsResponse{}
	_body, _err := client.ListUserActiveRunsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出工作空间
//
// @param request - ListWorkspacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithOptions(request *ListWorkspacesRequest, runtime *util.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkspaces"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出工作空间
//
// @param request - ListWorkspacesRequest
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 恢复投递
//
// @param request - ResumeSubmissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeSubmissionResponse
func (client *Client) ResumeSubmissionWithOptions(request *ResumeSubmissionRequest, runtime *util.RuntimeOptions) (_result *ResumeSubmissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubmissionId)) {
		query["SubmissionId"] = request.SubmissionId
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeSubmission"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeSubmissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 恢复投递
//
// @param request - ResumeSubmissionRequest
//
// @return ResumeSubmissionResponse
func (client *Client) ResumeSubmission(request *ResumeSubmissionRequest) (_result *ResumeSubmissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeSubmissionResponse{}
	_body, _err := client.ResumeSubmissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用版本打标
//
// @param request - TagAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagAppResponse
func (client *Client) TagAppWithOptions(request *TagAppRequest, runtime *util.RuntimeOptions) (_result *TagAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppRevision)) {
		query["AppRevision"] = request.AppRevision
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionTag)) {
		query["RevisionTag"] = request.RevisionTag
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagApp"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用版本打标
//
// @param request - TagAppRequest
//
// @return TagAppResponse
func (client *Client) TagApp(request *TagAppRequest) (_result *TagAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagAppResponse{}
	_body, _err := client.TagAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateEntity is deprecated, please use EasyGene::2021-03-15::UpdateEntityItems instead.
//
// Summary:
//
// 修改数据表格
//
// @param tmpReq - UpdateEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEntityResponse
// Deprecated
func (client *Client) UpdateEntityWithOptions(tmpReq *UpdateEntityRequest, runtime *util.RuntimeOptions) (_result *UpdateEntityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityItems)) {
		request.EntityItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityItems, tea.String("EntityItems"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityItemsShrink)) {
		body["EntityItems"] = request.EntityItemsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEntity"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateEntity is deprecated, please use EasyGene::2021-03-15::UpdateEntityItems instead.
//
// Summary:
//
// 修改数据表格
//
// @param request - UpdateEntityRequest
//
// @return UpdateEntityResponse
// Deprecated
func (client *Client) UpdateEntity(request *UpdateEntityRequest) (_result *UpdateEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEntityResponse{}
	_body, _err := client.UpdateEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改表格数据
//
// @param tmpReq - UpdateEntityItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEntityItemsResponse
func (client *Client) UpdateEntityItemsWithOptions(tmpReq *UpdateEntityItemsRequest, runtime *util.RuntimeOptions) (_result *UpdateEntityItemsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateEntityItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityItems)) {
		request.EntityItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityItems, tea.String("EntityItems"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityItemsShrink)) {
		body["EntityItems"] = request.EntityItemsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEntityItems"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEntityItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改表格数据
//
// @param request - UpdateEntityItemsRequest
//
// @return UpdateEntityItemsResponse
func (client *Client) UpdateEntityItems(request *UpdateEntityItemsRequest) (_result *UpdateEntityItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEntityItemsResponse{}
	_body, _err := client.UpdateEntityItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用模板
//
// @param tmpReq - UpdateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithOptions(tmpReq *UpdateTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InputsExpression)) {
		request.InputsExpressionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputsExpression, tea.String("InputsExpression"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OutputsExpression)) {
		request.OutputsExpressionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutputsExpression, tea.String("OutputsExpression"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.RootEntity)) {
		query["RootEntity"] = request.RootEntity
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InputsExpressionShrink)) {
		body["InputsExpression"] = request.InputsExpressionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OutputsExpressionShrink)) {
		body["OutputsExpression"] = request.OutputsExpressionShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTemplate"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用模板
//
// @param request - UpdateTemplateRequest
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplate(request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改工作空间
//
// @param request - UpdateWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceResponse
func (client *Client) UpdateWorkspaceWithOptions(request *UpdateWorkspaceRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.JobLifecycle)) {
		query["JobLifecycle"] = request.JobLifecycle
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkspace"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改工作空间
//
// @param request - UpdateWorkspaceRequest
//
// @return UpdateWorkspaceResponse
func (client *Client) UpdateWorkspace(request *UpdateWorkspaceRequest) (_result *UpdateWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.UpdateWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传数据表格
//
// @param request - UploadEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadEntityResponse
func (client *Client) UploadEntityWithOptions(request *UploadEntityRequest, runtime *util.RuntimeOptions) (_result *UploadEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityCSVFile)) {
		query["EntityCSVFile"] = request.EntityCSVFile
	}

	if !tea.BoolValue(util.IsUnset(request.Workspace)) {
		query["Workspace"] = request.Workspace
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadEntity"),
		Version:     tea.String("2021-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传数据表格
//
// @param request - UploadEntityRequest
//
// @return UploadEntityResponse
func (client *Client) UploadEntity(request *UploadEntityRequest) (_result *UploadEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadEntityResponse{}
	_body, _err := client.UploadEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
