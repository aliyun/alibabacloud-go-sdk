// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchModifyInstanceStatusRequest struct {
	// Specifies whether to start or stop the playbook.
	//
	// 	- **0**: stops the playbook.
	//
	// 	- **1**: starts the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Active *int32 `json:"Active,omitempty" xml:"Active,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The playbook UUID. If you want to specify multiple playbooks, separate the playbook UUIDs with commas (,).
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-1xxxxxx,s8df2e-s8dfs-xxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s BatchModifyInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *BatchModifyInstanceStatusRequest) SetActive(v int32) *BatchModifyInstanceStatusRequest {
	s.Active = &v
	return s
}

func (s *BatchModifyInstanceStatusRequest) SetLang(v string) *BatchModifyInstanceStatusRequest {
	s.Lang = &v
	return s
}

func (s *BatchModifyInstanceStatusRequest) SetPlaybookUuid(v string) *BatchModifyInstanceStatusRequest {
	s.PlaybookUuid = &v
	return s
}

type BatchModifyInstanceStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 358E012F-B516-599D-9ED0-A1A361CDE615
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchModifyInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *BatchModifyInstanceStatusResponseBody) SetRequestId(v string) *BatchModifyInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

type BatchModifyInstanceStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchModifyInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchModifyInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *BatchModifyInstanceStatusResponse) SetHeaders(v map[string]*string) *BatchModifyInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *BatchModifyInstanceStatusResponse) SetStatusCode(v int32) *BatchModifyInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchModifyInstanceStatusResponse) SetBody(v *BatchModifyInstanceStatusResponseBody) *BatchModifyInstanceStatusResponse {
	s.Body = v
	return s
}

type ComparePlaybooksRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the second version.
	//
	// >  You can call the [DescribePlaybookReleases](~~DescribePlaybookReleases~~) operation to query the IDs of versions. The system automatically generates IDs for new versions.
	//
	// This parameter is required.
	//
	// example:
	//
	// sfdf2395-e814-459f-9662-xxxxx
	NewPlaybookReleaseId *int32 `json:"NewPlaybookReleaseId,omitempty" xml:"NewPlaybookReleaseId,omitempty"`
	// The ID of the first version.
	//
	// >  You can call the [DescribePlaybookReleases](~~DescribePlaybookReleases~~) operation to query the IDs of versions. The system automatically generates IDs for new versions.
	//
	// This parameter is required.
	//
	// example:
	//
	// sflk23423-e814-459f-9662-xxxxx
	OldPlaybookReleaseId *int32 `json:"OldPlaybookReleaseId,omitempty" xml:"OldPlaybookReleaseId,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// f916b93e-e814-459f-9662-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s ComparePlaybooksRequest) String() string {
	return tea.Prettify(s)
}

func (s ComparePlaybooksRequest) GoString() string {
	return s.String()
}

func (s *ComparePlaybooksRequest) SetLang(v string) *ComparePlaybooksRequest {
	s.Lang = &v
	return s
}

func (s *ComparePlaybooksRequest) SetNewPlaybookReleaseId(v int32) *ComparePlaybooksRequest {
	s.NewPlaybookReleaseId = &v
	return s
}

func (s *ComparePlaybooksRequest) SetOldPlaybookReleaseId(v int32) *ComparePlaybooksRequest {
	s.OldPlaybookReleaseId = &v
	return s
}

func (s *ComparePlaybooksRequest) SetPlaybookUuid(v string) *ComparePlaybooksRequest {
	s.PlaybookUuid = &v
	return s
}

type ComparePlaybooksResponseBody struct {
	// The comparison result.
	CompareResult *ComparePlaybooksResponseBodyCompareResult `json:"CompareResult,omitempty" xml:"CompareResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2EC05B06-BF3C-5F3E-8FE8-3B1FAD76087A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ComparePlaybooksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ComparePlaybooksResponseBody) GoString() string {
	return s.String()
}

func (s *ComparePlaybooksResponseBody) SetCompareResult(v *ComparePlaybooksResponseBodyCompareResult) *ComparePlaybooksResponseBody {
	s.CompareResult = v
	return s
}

func (s *ComparePlaybooksResponseBody) SetRequestId(v string) *ComparePlaybooksResponseBody {
	s.RequestId = &v
	return s
}

type ComparePlaybooksResponseBodyCompareResult struct {
	// The description of the comparison result.
	//
	// example:
	//
	// The first version adds one node compared to the second version
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the second version provides more information than the first version. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	New *bool `json:"New,omitempty" xml:"New,omitempty"`
	// Indicates whether the configurations of the two versions are the same. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Same *bool `json:"Same,omitempty" xml:"Same,omitempty"`
}

func (s ComparePlaybooksResponseBodyCompareResult) String() string {
	return tea.Prettify(s)
}

func (s ComparePlaybooksResponseBodyCompareResult) GoString() string {
	return s.String()
}

func (s *ComparePlaybooksResponseBodyCompareResult) SetDescription(v string) *ComparePlaybooksResponseBodyCompareResult {
	s.Description = &v
	return s
}

func (s *ComparePlaybooksResponseBodyCompareResult) SetNew(v bool) *ComparePlaybooksResponseBodyCompareResult {
	s.New = &v
	return s
}

func (s *ComparePlaybooksResponseBodyCompareResult) SetSame(v bool) *ComparePlaybooksResponseBodyCompareResult {
	s.Same = &v
	return s
}

type ComparePlaybooksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ComparePlaybooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ComparePlaybooksResponse) String() string {
	return tea.Prettify(s)
}

func (s ComparePlaybooksResponse) GoString() string {
	return s.String()
}

func (s *ComparePlaybooksResponse) SetHeaders(v map[string]*string) *ComparePlaybooksResponse {
	s.Headers = v
	return s
}

func (s *ComparePlaybooksResponse) SetStatusCode(v int32) *ComparePlaybooksResponse {
	s.StatusCode = &v
	return s
}

func (s *ComparePlaybooksResponse) SetBody(v *ComparePlaybooksResponseBody) *ComparePlaybooksResponse {
	s.Body = v
	return s
}

type ConvertPlaybookRequest struct {
	// Language type for request and response messages. Values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// User ID for the administrator to switch to another member\\"s perspective.
	//
	// example:
	//
	// 13760*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// View type. Values:
	//
	// - 0: Current Alibaba Cloud account view.
	//
	// - 1: View for all accounts under the enterprise.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// XML configuration information for playbook orchestration.
	//
	// This parameter is required.
	//
	// example:
	//
	// <?xml version=\\"1.0\\" encoding=\\"UTF-8\\"?>
	//
	// <bpmn:definitions xmlns:xsi=\\"http://www.w3.org/2001/XMLSchema-instance\\" xmlns:bpmn=\\"http://www.omg.org/spec/BPMN/20100524/MODEL\\" xmlns:bpmndi=\\"http://www.omg.org/spec/BPMN/20100524/DI\\" xmlns:dc=\\"http://www.omg.org/spec/DD/20100524/DC\\" id=\\"Definitions_1\\" targetNamespace=\\"http://bpmn.io/schema/bpmn\\">
	//
	//   <bpmn:process id=\\"Process_1\\" isExecutable=\\"false\\">
	//
	//     <bpmn:startEvent id=\\"StartEvent_1\\" />
	//
	//   </bpmn:process>
	//
	//   <bpmndi:BPMNDiagram id=\\"BPMNDiagram_1\\">
	//
	//      <bpmndi:BPMNPlane id=\\"BPMNPlane_1\\" bpmnElement=\\"Process_1\\">
	//
	//            <bpmndi:BPMNShape id=\\"_BPMNShape_StartEvent_2\\" bpmnElement=\\"StartEvent_1\\">
	//
	//                    <dc:Bounds x=\\"173\\" y=\\"102\\" width=\\"36\\" height=\\"36\\" />
	//
	//             </bpmndi:BPMNShape>
	//
	//      </bpmndi:BPMNPlane>
	//
	//   </bpmndi:BPMNDiagram>
	//
	// </bpmn:definitions>
	Taskflow *string `json:"Taskflow,omitempty" xml:"Taskflow,omitempty"`
}

func (s ConvertPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertPlaybookRequest) GoString() string {
	return s.String()
}

func (s *ConvertPlaybookRequest) SetLang(v string) *ConvertPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *ConvertPlaybookRequest) SetRoleFor(v int64) *ConvertPlaybookRequest {
	s.RoleFor = &v
	return s
}

func (s *ConvertPlaybookRequest) SetRoleType(v string) *ConvertPlaybookRequest {
	s.RoleType = &v
	return s
}

func (s *ConvertPlaybookRequest) SetTaskflow(v string) *ConvertPlaybookRequest {
	s.Taskflow = &v
	return s
}

type ConvertPlaybookResponseBody struct {
	// The configurations.
	//
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of this call request, which is a unique identifier generated by Alibaba Cloud for this request, and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// 39C38A34-****-*****-****-7263B435C316
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertPlaybookResponseBody) SetConfig(v string) *ConvertPlaybookResponseBody {
	s.Config = &v
	return s
}

func (s *ConvertPlaybookResponseBody) SetRequestId(v string) *ConvertPlaybookResponseBody {
	s.RequestId = &v
	return s
}

type ConvertPlaybookResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertPlaybookResponse) GoString() string {
	return s.String()
}

func (s *ConvertPlaybookResponse) SetHeaders(v map[string]*string) *ConvertPlaybookResponse {
	s.Headers = v
	return s
}

func (s *ConvertPlaybookResponse) SetStatusCode(v int32) *ConvertPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertPlaybookResponse) SetBody(v *ConvertPlaybookResponseBody) *ConvertPlaybookResponse {
	s.Body = v
	return s
}

type CopyPlaybookRequest struct {
	// example:
	//
	// playbook description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// playbook_xxx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 0
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// example:
	//
	// 137602*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 94bc318c-****-4cba-****-801ccb0d739f
	SourcePlaybookUuid *string `json:"SourcePlaybookUuid,omitempty" xml:"SourcePlaybookUuid,omitempty"`
}

func (s CopyPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyPlaybookRequest) GoString() string {
	return s.String()
}

func (s *CopyPlaybookRequest) SetDescription(v string) *CopyPlaybookRequest {
	s.Description = &v
	return s
}

func (s *CopyPlaybookRequest) SetDisplayName(v string) *CopyPlaybookRequest {
	s.DisplayName = &v
	return s
}

func (s *CopyPlaybookRequest) SetLang(v string) *CopyPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *CopyPlaybookRequest) SetReleaseVersion(v string) *CopyPlaybookRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *CopyPlaybookRequest) SetRoleFor(v int64) *CopyPlaybookRequest {
	s.RoleFor = &v
	return s
}

func (s *CopyPlaybookRequest) SetRoleType(v string) *CopyPlaybookRequest {
	s.RoleType = &v
	return s
}

func (s *CopyPlaybookRequest) SetSourcePlaybookUuid(v string) *CopyPlaybookRequest {
	s.SourcePlaybookUuid = &v
	return s
}

type CopyPlaybookResponseBody struct {
	Data *CopyPlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Page *CopyPlaybookResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 2EC05B06-****-5F3E-****-3B1FAD76087A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *CopyPlaybookResponseBody) SetData(v *CopyPlaybookResponseBodyData) *CopyPlaybookResponseBody {
	s.Data = v
	return s
}

func (s *CopyPlaybookResponseBody) SetPage(v *CopyPlaybookResponseBodyPage) *CopyPlaybookResponseBody {
	s.Page = v
	return s
}

func (s *CopyPlaybookResponseBody) SetRequestId(v string) *CopyPlaybookResponseBody {
	s.RequestId = &v
	return s
}

type CopyPlaybookResponseBodyData struct {
	// example:
	//
	// 1
	Active *int32 `json:"Active,omitempty" xml:"Active,omitempty"`
	// example:
	//
	// This is a action of processing for WAF
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 11111
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 1
	FailNum *int32 `json:"FailNum,omitempty" xml:"FailNum,omitempty"`
	// example:
	//
	// 0.5
	FailRate *float64 `json:"FailRate,omitempty" xml:"FailRate,omitempty"`
	// example:
	//
	// 1655951601000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1638270967000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1
	HistoryMd5 *int32 `json:"HistoryMd5,omitempty" xml:"HistoryMd5,omitempty"`
	// example:
	//
	// [{\\"name\\":\\"1\\",\\"dataType\\":\\"String\\",\\"required\\":false,\\"isArray\\":false,\\"example\\":\\"\\",\\"description\\":\\"\\",\\"id\\":0,\\"typeName\\":\\"String\\",\\"dataClass\\":\\"normal\\"}]
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// example:
	//
	// 1725258397847
	LastRuntime *int64 `json:"LastRuntime,omitempty" xml:"LastRuntime,omitempty"`
	// example:
	//
	// 037046****1b00c4717963818ccbf2xx
	LogicReleaseTaskflowMd5 *string `json:"LogicReleaseTaskflowMd5,omitempty" xml:"LogicReleaseTaskflowMd5,omitempty"`
	// example:
	//
	// []
	OutputParams *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	// example:
	//
	// user
	OwnType *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// example:
	//
	// 1
	Permission *int32 `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// example:
	//
	// 1
	PlaybookStatus *int32 `json:"PlaybookStatus,omitempty" xml:"PlaybookStatus,omitempty"`
	// example:
	//
	// 9e38111e-9794-4784-9ca8-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// example:
	//
	// 1
	SuccNum *int32 `json:"SuccNum,omitempty" xml:"SuccNum,omitempty"`
	// example:
	//
	// 13760*****8718726
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CopyPlaybookResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CopyPlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *CopyPlaybookResponseBodyData) SetActive(v int32) *CopyPlaybookResponseBodyData {
	s.Active = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetDescription(v string) *CopyPlaybookResponseBodyData {
	s.Description = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetDisplayName(v string) *CopyPlaybookResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetFailNum(v int32) *CopyPlaybookResponseBodyData {
	s.FailNum = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetFailRate(v float64) *CopyPlaybookResponseBodyData {
	s.FailRate = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetGmtCreate(v int64) *CopyPlaybookResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetGmtModified(v int64) *CopyPlaybookResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetHistoryMd5(v int32) *CopyPlaybookResponseBodyData {
	s.HistoryMd5 = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetInputParams(v string) *CopyPlaybookResponseBodyData {
	s.InputParams = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetLastRuntime(v int64) *CopyPlaybookResponseBodyData {
	s.LastRuntime = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetLogicReleaseTaskflowMd5(v string) *CopyPlaybookResponseBodyData {
	s.LogicReleaseTaskflowMd5 = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetOutputParams(v string) *CopyPlaybookResponseBodyData {
	s.OutputParams = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetOwnType(v string) *CopyPlaybookResponseBodyData {
	s.OwnType = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetPermission(v int32) *CopyPlaybookResponseBodyData {
	s.Permission = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetPlaybookStatus(v int32) *CopyPlaybookResponseBodyData {
	s.PlaybookStatus = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetPlaybookUuid(v string) *CopyPlaybookResponseBodyData {
	s.PlaybookUuid = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetSuccNum(v int32) *CopyPlaybookResponseBodyData {
	s.SuccNum = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetTenantId(v string) *CopyPlaybookResponseBodyData {
	s.TenantId = &v
	return s
}

type CopyPlaybookResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CopyPlaybookResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s CopyPlaybookResponseBodyPage) GoString() string {
	return s.String()
}

func (s *CopyPlaybookResponseBodyPage) SetPageNumber(v int32) *CopyPlaybookResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *CopyPlaybookResponseBodyPage) SetPageSize(v int32) *CopyPlaybookResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *CopyPlaybookResponseBodyPage) SetTotalCount(v int32) *CopyPlaybookResponseBodyPage {
	s.TotalCount = &v
	return s
}

type CopyPlaybookResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyPlaybookResponse) GoString() string {
	return s.String()
}

func (s *CopyPlaybookResponse) SetHeaders(v map[string]*string) *CopyPlaybookResponse {
	s.Headers = v
	return s
}

func (s *CopyPlaybookResponse) SetStatusCode(v int32) *CopyPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyPlaybookResponse) SetBody(v *CopyPlaybookResponseBody) *CopyPlaybookResponse {
	s.Body = v
	return s
}

type CreatePlaybookRequest struct {
	// Description of the playbook.
	//
	// example:
	//
	// This is a new version
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// test09
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Language type for receiving messages. Values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Playbook TaskFlow type.
	//
	// - **x6*	- : x6
	//
	// - **bpmn**: bpmn
	//
	// example:
	//
	// x6
	TaskflowType *string `json:"TaskflowType,omitempty" xml:"TaskflowType,omitempty"`
}

func (s CreatePlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePlaybookRequest) GoString() string {
	return s.String()
}

func (s *CreatePlaybookRequest) SetDescription(v string) *CreatePlaybookRequest {
	s.Description = &v
	return s
}

func (s *CreatePlaybookRequest) SetDisplayName(v string) *CreatePlaybookRequest {
	s.DisplayName = &v
	return s
}

func (s *CreatePlaybookRequest) SetLang(v string) *CreatePlaybookRequest {
	s.Lang = &v
	return s
}

func (s *CreatePlaybookRequest) SetTaskflowType(v string) *CreatePlaybookRequest {
	s.TaskflowType = &v
	return s
}

type CreatePlaybookResponseBody struct {
	// The result of the creation.
	Data *CreatePlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of this call request, a unique identifier generated by Alibaba Cloud for this request, which can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// B09B40B2-F11E-512C-B755-423F2056C17B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePlaybookResponseBody) SetData(v *CreatePlaybookResponseBodyData) *CreatePlaybookResponseBody {
	s.Data = v
	return s
}

func (s *CreatePlaybookResponseBody) SetRequestId(v string) *CreatePlaybookResponseBody {
	s.RequestId = &v
	return s
}

type CreatePlaybookResponseBodyData struct {
	// UUID of the newly created playbook.
	//
	// example:
	//
	// 9e38111e-9794-4784-9ca8-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s CreatePlaybookResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePlaybookResponseBodyData) SetPlaybookUuid(v string) *CreatePlaybookResponseBodyData {
	s.PlaybookUuid = &v
	return s
}

type CreatePlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePlaybookResponse) GoString() string {
	return s.String()
}

func (s *CreatePlaybookResponse) SetHeaders(v map[string]*string) *CreatePlaybookResponse {
	s.Headers = v
	return s
}

func (s *CreatePlaybookResponse) SetStatusCode(v int32) *CreatePlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePlaybookResponse) SetBody(v *CreatePlaybookResponseBody) *CreatePlaybookResponse {
	s.Body = v
	return s
}

type DebugPlaybookRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f916b93e-e814-459f-9662-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The input parameters that you use to debug the playbook. You can define the parameters based on your business requirements.
	//
	// example:
	//
	// {
	//
	//    "param1":"a",
	//
	//    "param2":"b"
	//
	// }
	Record *string `json:"Record,omitempty" xml:"Record,omitempty"`
	// The XML configuration of the playbook.
	//
	// >  You can call the [DescribePlaybook](~~DescribePlaybook~~) operation to query the XML configuration of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// <?xml version="1.0" encoding="UTF-8"?><bpmn:definitions xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:bpmn="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:dc="http://www.omg.org/spec/DD/20100524/DC" targetNamespace="http://bpmn.io/schema/bpmn" id="Definitions_1"><bpmn:process id="Process_1" isExecutable="false"><bpmn:startEvent id="StartEvent_1"/></bpmn:process><bpmndi:BPMNDiagram id="BPMNDiagram_1"><bpmndi:BPMNPlane id="BPMNPlane_1" bpmnElement="Process_1"><bpmndi:BPMNShape id="_BPMNShape_StartEvent_2" bpmnElement="StartEvent_1"><dc:Bounds height="36.0" width="36.0" x="173.0" y="102.0"/></bpmndi:BPMNShape></bpmndi:BPMNPlane></bpmndi:BPMNDiagram></bpmn:definitions>
	Taskflow *string `json:"Taskflow,omitempty" xml:"Taskflow,omitempty"`
}

func (s DebugPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s DebugPlaybookRequest) GoString() string {
	return s.String()
}

func (s *DebugPlaybookRequest) SetLang(v string) *DebugPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DebugPlaybookRequest) SetPlaybookUuid(v string) *DebugPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DebugPlaybookRequest) SetRecord(v string) *DebugPlaybookRequest {
	s.Record = &v
	return s
}

func (s *DebugPlaybookRequest) SetTaskflow(v string) *DebugPlaybookRequest {
	s.Taskflow = &v
	return s
}

type DebugPlaybookResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 75E56B2C-C8FA-5A2F-AA08-8745E2AC33EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The UUID of the debugging task. You can use the UUID to query the result and other details of the debugging task.
	//
	// example:
	//
	// 6d412cfa-0905-4567-8a83-xxxxxx
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
}

func (s DebugPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DebugPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DebugPlaybookResponseBody) SetRequestId(v string) *DebugPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugPlaybookResponseBody) SetRequestUuid(v string) *DebugPlaybookResponseBody {
	s.RequestUuid = &v
	return s
}

type DebugPlaybookResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DebugPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s DebugPlaybookResponse) GoString() string {
	return s.String()
}

func (s *DebugPlaybookResponse) SetHeaders(v map[string]*string) *DebugPlaybookResponse {
	s.Headers = v
	return s
}

func (s *DebugPlaybookResponse) SetStatusCode(v int32) *DebugPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugPlaybookResponse) SetBody(v *DebugPlaybookResponseBody) *DebugPlaybookResponse {
	s.Body = v
	return s
}

type DeleteComponentAssetRequest struct {
	// The ID of the asset.
	//
	// >  You can call the [DescribeComponentAssets](~~DescribeComponentAssets~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12x
	AssetId *int64 `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DeleteComponentAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteComponentAssetRequest) GoString() string {
	return s.String()
}

func (s *DeleteComponentAssetRequest) SetAssetId(v int64) *DeleteComponentAssetRequest {
	s.AssetId = &v
	return s
}

func (s *DeleteComponentAssetRequest) SetLang(v string) *DeleteComponentAssetRequest {
	s.Lang = &v
	return s
}

type DeleteComponentAssetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 39C38A34-8532-5D44-B88A-7263B435C316
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteComponentAssetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteComponentAssetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComponentAssetResponseBody) SetRequestId(v string) *DeleteComponentAssetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteComponentAssetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComponentAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComponentAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteComponentAssetResponse) GoString() string {
	return s.String()
}

func (s *DeleteComponentAssetResponse) SetHeaders(v map[string]*string) *DeleteComponentAssetResponse {
	s.Headers = v
	return s
}

func (s *DeleteComponentAssetResponse) SetStatusCode(v int32) *DeleteComponentAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComponentAssetResponse) SetBody(v *DeleteComponentAssetResponseBody) *DeleteComponentAssetResponse {
	s.Body = v
	return s
}

type DeletePlaybookRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e99dab31-499b-4307-9248-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DeletePlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePlaybookRequest) GoString() string {
	return s.String()
}

func (s *DeletePlaybookRequest) SetLang(v string) *DeletePlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DeletePlaybookRequest) SetPlaybookUuid(v string) *DeletePlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

type DeletePlaybookResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6F3CA8A9-B5BB-506A-9182-FFE80A6E0584
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePlaybookResponseBody) SetRequestId(v string) *DeletePlaybookResponseBody {
	s.RequestId = &v
	return s
}

type DeletePlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePlaybookResponse) GoString() string {
	return s.String()
}

func (s *DeletePlaybookResponse) SetHeaders(v map[string]*string) *DeletePlaybookResponse {
	s.Headers = v
	return s
}

func (s *DeletePlaybookResponse) SetStatusCode(v int32) *DeletePlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePlaybookResponse) SetBody(v *DeletePlaybookResponseBody) *DeletePlaybookResponse {
	s.Body = v
	return s
}

type DescribeComponentAssetFormRequest struct {
	// The component name.
	//
	// This parameter is required.
	//
	// example:
	//
	// python3
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeComponentAssetFormRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetFormRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetFormRequest) SetComponentName(v string) *DescribeComponentAssetFormRequest {
	s.ComponentName = &v
	return s
}

func (s *DescribeComponentAssetFormRequest) SetLang(v string) *DescribeComponentAssetFormRequest {
	s.Lang = &v
	return s
}

type DescribeComponentAssetFormResponseBody struct {
	// The metadata of the asset in the component. The value is a JSON array and contains the following fields:
	//
	// 	- **name**: the parameter name.
	//
	// 	- **defaultValue**: the default parameter value.
	//
	// 	- **description**: the parameter description.
	//
	// 	- **required**: indicates whether the parameter is required. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "defaultValue": "",
	//
	//         "description": "assetname",
	//
	//         "name": "assetname",
	//
	//         "required": true
	//
	//     }
	//
	// ]
	ComponentAssetForm *string `json:"ComponentAssetForm,omitempty" xml:"ComponentAssetForm,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9D1651AC-31CC-5CC4-A14E-626B3FCC1022
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentAssetFormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetFormResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetFormResponseBody) SetComponentAssetForm(v string) *DescribeComponentAssetFormResponseBody {
	s.ComponentAssetForm = &v
	return s
}

func (s *DescribeComponentAssetFormResponseBody) SetRequestId(v string) *DescribeComponentAssetFormResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComponentAssetFormResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentAssetFormResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentAssetFormResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetFormResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetFormResponse) SetHeaders(v map[string]*string) *DescribeComponentAssetFormResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentAssetFormResponse) SetStatusCode(v int32) *DescribeComponentAssetFormResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentAssetFormResponse) SetBody(v *DescribeComponentAssetFormResponseBody) *DescribeComponentAssetFormResponse {
	s.Body = v
	return s
}

type DescribeComponentAssetsRequest struct {
	// The name of the component.
	//
	// This parameter is required.
	//
	// example:
	//
	// python3
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeComponentAssetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetsRequest) SetComponentName(v string) *DescribeComponentAssetsRequest {
	s.ComponentName = &v
	return s
}

func (s *DescribeComponentAssetsRequest) SetLang(v string) *DescribeComponentAssetsRequest {
	s.Lang = &v
	return s
}

type DescribeComponentAssetsResponseBody struct {
	// The information about the assets.
	ComponentAssets []*DescribeComponentAssetsResponseBodyComponentAssets `json:"ComponentAssets,omitempty" xml:"ComponentAssets,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BFEFB76D-DD0E-5529-BD57-0DAC10B9B30F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentAssetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetsResponseBody) SetComponentAssets(v []*DescribeComponentAssetsResponseBodyComponentAssets) *DescribeComponentAssetsResponseBody {
	s.ComponentAssets = v
	return s
}

func (s *DescribeComponentAssetsResponseBody) SetRequestId(v string) *DescribeComponentAssetsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComponentAssetsResponseBodyComponentAssets struct {
	// The UUID of the asset.
	//
	// example:
	//
	// ff6fe161-93e2-464c-a326-fxxxxxx
	AssetUuid *string `json:"AssetUuid,omitempty" xml:"AssetUuid,omitempty"`
	// The name of the component to which the asset belongs.
	//
	// example:
	//
	// pyhton3
	Componentname *string `json:"Componentname,omitempty" xml:"Componentname,omitempty"`
	// The time when the asset was created. The time is in the yyyy-MM-ddTHH:mm:ssZ format and is displayed in UTC.
	//
	// example:
	//
	// 2023-03-23T14:38Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the asset was modified. The time is in the yyyy-MM-ddTHH:mm:ssZ format and is displayed in UTC.
	//
	// example:
	//
	// 2023-03-23T14:38Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The UUID of the asset.
	//
	// example:
	//
	// 7xx
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// test asset
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of the asset in the JSON string format. DescribeComponentAssetForm
	//
	// >  For more information, see [DescribeComponentAssetForm](~~DescribeComponentAssetForm~~).
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "name": "authMethod",
	//
	//         "value": "ak"
	//
	//     },
	//
	//     {
	//
	//         "name": "accessKeyId",
	//
	//         "value": "xxxxxxx"
	//
	//     },
	//
	//     {
	//
	//         "name": "accessKeySecret",
	//
	//         "value": "xxxxx"
	//
	//     },
	//
	//     {
	//
	//         "name": "roleArn",
	//
	//         "value": ""
	//
	//     }
	//
	// ]
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s DescribeComponentAssetsResponseBodyComponentAssets) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetsResponseBodyComponentAssets) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetAssetUuid(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.AssetUuid = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetComponentname(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.Componentname = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetGmtCreate(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.GmtCreate = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetGmtModified(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.GmtModified = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetId(v int64) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.Id = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetName(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.Name = &v
	return s
}

func (s *DescribeComponentAssetsResponseBodyComponentAssets) SetParams(v string) *DescribeComponentAssetsResponseBodyComponentAssets {
	s.Params = &v
	return s
}

type DescribeComponentAssetsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentAssetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentAssetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetsResponse) SetHeaders(v map[string]*string) *DescribeComponentAssetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentAssetsResponse) SetStatusCode(v int32) *DescribeComponentAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentAssetsResponse) SetBody(v *DescribeComponentAssetsResponseBody) *DescribeComponentAssetsResponse {
	s.Body = v
	return s
}

type DescribeComponentListRequest struct {
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// b724d2b0-3c3b-4223-9bfd-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeComponentListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentListRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentListRequest) SetLang(v string) *DescribeComponentListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeComponentListRequest) SetPlaybookUuid(v string) *DescribeComponentListRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribeComponentListResponseBody struct {
	// The information about the components. The value is a JSON array.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "actions": [
	//
	//             {
	//
	//                 "description": "mysql component",
	//
	//                 "name": "storeIdb",
	//
	//                 "parameters": [
	//
	//                     {
	//
	//                         "description": "update the mysql db",
	//
	//                         "name": "updateSql",
	//
	//                         "required": false
	//
	//                     }
	//
	//                 ]
	//
	//             }
	//
	//         ],
	//
	//         "basic": {
	//
	//             "description": "mysq sql component for 5.6",
	//
	//             "logo": "https://img.alicdn.com/tfs/TB1H89IpH3nBKNjSZFMXXaUSFXa-200-200.svg",
	//
	//             "name": "Mysql"
	//
	//         }
	//
	//     }
	//
	// ]
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B0A255B3-495C-56FB-8B6B-DB073F80388A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentListResponseBody) SetComponents(v string) *DescribeComponentListResponseBody {
	s.Components = &v
	return s
}

func (s *DescribeComponentListResponseBody) SetRequestId(v string) *DescribeComponentListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComponentListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentListResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentListResponse) SetHeaders(v map[string]*string) *DescribeComponentListResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentListResponse) SetStatusCode(v int32) *DescribeComponentListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentListResponse) SetBody(v *DescribeComponentListResponseBody) *DescribeComponentListResponse {
	s.Body = v
	return s
}

type DescribeComponentPlaybookRequest struct {
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac343acc-1a61-4084-9a1cxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeComponentPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentPlaybookRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentPlaybookRequest) SetLang(v string) *DescribeComponentPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DescribeComponentPlaybookRequest) SetPlaybookUuid(v string) *DescribeComponentPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribeComponentPlaybookResponseBody struct {
	// The information about the predefined components.
	Playbooks []*DescribeComponentPlaybookResponseBodyPlaybooks `json:"Playbooks,omitempty" xml:"Playbooks,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C5F5D6C9-DF1A-5381-92B1-39676F777D20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentPlaybookResponseBody) SetPlaybooks(v []*DescribeComponentPlaybookResponseBodyPlaybooks) *DescribeComponentPlaybookResponseBody {
	s.Playbooks = v
	return s
}

func (s *DescribeComponentPlaybookResponseBody) SetRequestId(v string) *DescribeComponentPlaybookResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComponentPlaybookResponseBodyPlaybooks struct {
	// The description of the predefined component.
	//
	// example:
	//
	// aegis_kill_process
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the predefined component.
	//
	// example:
	//
	// AegisKillQuara
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The input parameter configuration of the playbook. The value is a JSON array.
	//
	// >  For more information, see [DescribePlaybookInputOutput](~~DescribePlaybookInputOutput~~).
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "typeName": "String",
	//
	//         "dataClass": "normal",
	//
	//         "dataType": "String",
	//
	//         "description": "period",
	//
	//         "example": "",
	//
	//         "name": "period",
	//
	//         "required": false
	//
	//     }
	//
	// ]
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
}

func (s DescribeComponentPlaybookResponseBodyPlaybooks) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentPlaybookResponseBodyPlaybooks) GoString() string {
	return s.String()
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) SetDescription(v string) *DescribeComponentPlaybookResponseBodyPlaybooks {
	s.Description = &v
	return s
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) SetDisplayName(v string) *DescribeComponentPlaybookResponseBodyPlaybooks {
	s.DisplayName = &v
	return s
}

func (s *DescribeComponentPlaybookResponseBodyPlaybooks) SetInputParams(v string) *DescribeComponentPlaybookResponseBodyPlaybooks {
	s.InputParams = &v
	return s
}

type DescribeComponentPlaybookResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentPlaybookResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentPlaybookResponse) SetHeaders(v map[string]*string) *DescribeComponentPlaybookResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentPlaybookResponse) SetStatusCode(v int32) *DescribeComponentPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentPlaybookResponse) SetBody(v *DescribeComponentPlaybookResponseBody) *DescribeComponentPlaybookResponse {
	s.Body = v
	return s
}

type DescribeComponentsJsRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeComponentsJsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentsJsRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentsJsRequest) SetLang(v string) *DescribeComponentsJsRequest {
	s.Lang = &v
	return s
}

type DescribeComponentsJsResponseBody struct {
	// The configuration of the JavaScript file for the component.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "js": "https://xxxxx.oss-cn-zhangjiakou.aliyuncs.com/componentUpload/xxxxx",
	//
	//         "name": "python3",
	//
	//         "ownType": "sys"
	//
	//     }
	//
	// ]
	ComponentsJs *string `json:"ComponentsJs,omitempty" xml:"ComponentsJs,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 58A518BC-E4A8-5BD7-AFEA-366046ED9073
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentsJsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentsJsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentsJsResponseBody) SetComponentsJs(v string) *DescribeComponentsJsResponseBody {
	s.ComponentsJs = &v
	return s
}

func (s *DescribeComponentsJsResponseBody) SetRequestId(v string) *DescribeComponentsJsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComponentsJsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentsJsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentsJsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentsJsResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentsJsResponse) SetHeaders(v map[string]*string) *DescribeComponentsJsResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentsJsResponse) SetStatusCode(v int32) *DescribeComponentsJsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentsJsResponse) SetBody(v *DescribeComponentsJsResponseBody) *DescribeComponentsJsResponse {
	s.Body = v
	return s
}

type DescribeDistinctReleasesRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bc0b8424-535c-4ed5-bd94-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The MD5 value of the playbook XML configuration.
	//
	// example:
	//
	// be0a4ef084dd174abe47xxxxx
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
}

func (s DescribeDistinctReleasesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistinctReleasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDistinctReleasesRequest) SetLang(v string) *DescribeDistinctReleasesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDistinctReleasesRequest) SetPlaybookUuid(v string) *DescribeDistinctReleasesRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribeDistinctReleasesRequest) SetTaskflowMd5(v string) *DescribeDistinctReleasesRequest {
	s.TaskflowMd5 = &v
	return s
}

type DescribeDistinctReleasesResponseBody struct {
	// The information about versions.
	Records []*DescribeDistinctReleasesResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 145CACF6-D276-5197-8549-CB1AD76E2AC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDistinctReleasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistinctReleasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDistinctReleasesResponseBody) SetRecords(v []*DescribeDistinctReleasesResponseBodyRecords) *DescribeDistinctReleasesResponseBody {
	s.Records = v
	return s
}

func (s *DescribeDistinctReleasesResponseBody) SetRequestId(v string) *DescribeDistinctReleasesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDistinctReleasesResponseBodyRecords struct {
	// The version description.
	//
	// example:
	//
	// demo version
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The MD5 value of the version XML configuration.
	//
	// example:
	//
	// 17cf53049bc8efa941207xxxxx
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
	// The format of the playbook. Valid values:
	//
	// 	- **xml**: XML format.
	//
	// 	- **x6**: JSON format.
	//
	// example:
	//
	// x6
	TaskflowType *string `json:"TaskflowType,omitempty" xml:"TaskflowType,omitempty"`
}

func (s DescribeDistinctReleasesResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistinctReleasesResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDistinctReleasesResponseBodyRecords) SetDescription(v string) *DescribeDistinctReleasesResponseBodyRecords {
	s.Description = &v
	return s
}

func (s *DescribeDistinctReleasesResponseBodyRecords) SetTaskflowMd5(v string) *DescribeDistinctReleasesResponseBodyRecords {
	s.TaskflowMd5 = &v
	return s
}

func (s *DescribeDistinctReleasesResponseBodyRecords) SetTaskflowType(v string) *DescribeDistinctReleasesResponseBodyRecords {
	s.TaskflowType = &v
	return s
}

type DescribeDistinctReleasesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDistinctReleasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDistinctReleasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistinctReleasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDistinctReleasesResponse) SetHeaders(v map[string]*string) *DescribeDistinctReleasesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDistinctReleasesResponse) SetStatusCode(v int32) *DescribeDistinctReleasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDistinctReleasesResponse) SetBody(v *DescribeDistinctReleasesResponseBody) *DescribeDistinctReleasesResponse {
	s.Body = v
	return s
}

type DescribeEnumItemsRequest struct {
	// The type of the enumeration item. Valid values:
	//
	// 	- **process**: scenarios
	//
	// This parameter is required.
	//
	// example:
	//
	// process
	EnumType *string `json:"EnumType,omitempty" xml:"EnumType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese (default)
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeEnumItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnumItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnumItemsRequest) SetEnumType(v string) *DescribeEnumItemsRequest {
	s.EnumType = &v
	return s
}

func (s *DescribeEnumItemsRequest) SetLang(v string) *DescribeEnumItemsRequest {
	s.Lang = &v
	return s
}

type DescribeEnumItemsResponseBody struct {
	// The information about the enumeration item.
	Data []*DescribeEnumItemsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E7698CFB-4E1C-5840-8EC9-691B86729E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnumItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnumItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnumItemsResponseBody) SetData(v []*DescribeEnumItemsResponseBodyData) *DescribeEnumItemsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEnumItemsResponseBody) SetRequestId(v string) *DescribeEnumItemsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEnumItemsResponseBodyData struct {
	// The key of the enumeration item.
	//
	// example:
	//
	// system_xxxxx_process_book
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the enumeration item.
	//
	// example:
	//
	// system_xxxxx_process_book
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEnumItemsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnumItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEnumItemsResponseBodyData) SetKey(v string) *DescribeEnumItemsResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeEnumItemsResponseBodyData) SetValue(v string) *DescribeEnumItemsResponseBodyData {
	s.Value = &v
	return s
}

type DescribeEnumItemsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnumItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnumItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEnumItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnumItemsResponse) SetHeaders(v map[string]*string) *DescribeEnumItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnumItemsResponse) SetStatusCode(v int32) *DescribeEnumItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnumItemsResponse) SetBody(v *DescribeEnumItemsResponseBody) *DescribeEnumItemsResponse {
	s.Body = v
	return s
}

type DescribeExecutePlaybooksRequest struct {
	// The entity type of the script input parameter. When you want to query multiple entity types, separate them with commas.
	//
	// - **ip**: IP entity.
	//
	// - **file**: file entity.
	//
	// - **process**: process entity.
	//
	// - **incident**: incident entity.
	//
	// example:
	//
	// ip,file,process,host
	InputMode *string `json:"InputMode,omitempty" xml:"InputMode,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The input parameter type of the playbook.
	//
	// 	- **template-ip**
	//
	// 	- **template-file**
	//
	// 	- **template-process**
	//
	// 	- **custom**
	//
	// example:
	//
	// custom
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The playbook name. Fuzzy search is supported.
	//
	// example:
	//
	// demo_test
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the playbook UUID.
	//
	// example:
	//
	// f916b93e-e814-459f-9662-xxxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExecutePlaybooksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutePlaybooksRequest) GoString() string {
	return s.String()
}

func (s *DescribeExecutePlaybooksRequest) SetInputMode(v string) *DescribeExecutePlaybooksRequest {
	s.InputMode = &v
	return s
}

func (s *DescribeExecutePlaybooksRequest) SetLang(v string) *DescribeExecutePlaybooksRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExecutePlaybooksRequest) SetParamType(v string) *DescribeExecutePlaybooksRequest {
	s.ParamType = &v
	return s
}

func (s *DescribeExecutePlaybooksRequest) SetPlaybookName(v string) *DescribeExecutePlaybooksRequest {
	s.PlaybookName = &v
	return s
}

func (s *DescribeExecutePlaybooksRequest) SetUuid(v string) *DescribeExecutePlaybooksRequest {
	s.Uuid = &v
	return s
}

type DescribeExecutePlaybooksResponseBody struct {
	// The playbook.
	PlaybookMetrics []*DescribeExecutePlaybooksResponseBodyPlaybookMetrics `json:"PlaybookMetrics,omitempty" xml:"PlaybookMetrics,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 88A39217-2802-5B1E-BA2B-CF1BBC43C1F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExecutePlaybooksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutePlaybooksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExecutePlaybooksResponseBody) SetPlaybookMetrics(v []*DescribeExecutePlaybooksResponseBodyPlaybookMetrics) *DescribeExecutePlaybooksResponseBody {
	s.PlaybookMetrics = v
	return s
}

func (s *DescribeExecutePlaybooksResponseBody) SetRequestId(v string) *DescribeExecutePlaybooksResponseBody {
	s.RequestId = &v
	return s
}

type DescribeExecutePlaybooksResponseBodyPlaybookMetrics struct {
	// The playbook description.
	//
	// example:
	//
	// a demo playbook
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The playbook name.
	//
	// example:
	//
	// demo_playbook
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The configuration of the input parameter. The value is a JSON array.
	//
	// >  For more information, see [DescribePlaybookInputOutput](~~DescribePlaybookInputOutput~~).
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "typeName": "String",
	//
	//         "dataClass": "normal",
	//
	//         "dataType": "String",
	//
	//         "description": "period",
	//
	//         "example": "",
	//
	//         "name": "period",
	//
	//         "required": false
	//
	//     }
	//
	// ]
	ParamConfig *string `json:"ParamConfig,omitempty" xml:"ParamConfig,omitempty"`
	// The input parameter type of the playbook.
	//
	// 	- **template-ip**
	//
	// 	- **template-file**
	//
	// 	- **template-process**
	//
	// 	- **custom**
	//
	// example:
	//
	// custom
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The playbook UUID.
	//
	// example:
	//
	// c5c88b5e-97ca-435d-8c20-2xxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExecutePlaybooksResponseBodyPlaybookMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutePlaybooksResponseBodyPlaybookMetrics) GoString() string {
	return s.String()
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetDescription(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.Description = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetDisplayName(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.DisplayName = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetParamConfig(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.ParamConfig = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetParamType(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.ParamType = &v
	return s
}

func (s *DescribeExecutePlaybooksResponseBodyPlaybookMetrics) SetUuid(v string) *DescribeExecutePlaybooksResponseBodyPlaybookMetrics {
	s.Uuid = &v
	return s
}

type DescribeExecutePlaybooksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExecutePlaybooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExecutePlaybooksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutePlaybooksResponse) GoString() string {
	return s.String()
}

func (s *DescribeExecutePlaybooksResponse) SetHeaders(v map[string]*string) *DescribeExecutePlaybooksResponse {
	s.Headers = v
	return s
}

func (s *DescribeExecutePlaybooksResponse) SetStatusCode(v int32) *DescribeExecutePlaybooksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExecutePlaybooksResponse) SetBody(v *DescribeExecutePlaybooksResponseBody) *DescribeExecutePlaybooksResponse {
	s.Body = v
	return s
}

type DescribeFieldRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The key of the global configuration. Valid values:
	//
	// 	- **soar_filed_tags**: queries the input template of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// soar_filed_tags
	QueryKey *string `json:"QueryKey,omitempty" xml:"QueryKey,omitempty"`
}

func (s DescribeFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFieldRequest) GoString() string {
	return s.String()
}

func (s *DescribeFieldRequest) SetLang(v string) *DescribeFieldRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFieldRequest) SetQueryKey(v string) *DescribeFieldRequest {
	s.QueryKey = &v
	return s
}

type DescribeFieldResponseBody struct {
	// The configuration content.
	//
	// example:
	//
	// ["ip","name","hostinfo","md5"]
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// The name of the global configuration.
	//
	// example:
	//
	// soar_filed_tags
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BCDE6498-83CC-50A1-8307-3D5A539C42F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFieldResponseBody) SetFields(v string) *DescribeFieldResponseBody {
	s.Fields = &v
	return s
}

func (s *DescribeFieldResponseBody) SetName(v string) *DescribeFieldResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFieldResponseBody) SetRequestId(v string) *DescribeFieldResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFieldResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFieldResponse) GoString() string {
	return s.String()
}

func (s *DescribeFieldResponse) SetHeaders(v map[string]*string) *DescribeFieldResponse {
	s.Headers = v
	return s
}

func (s *DescribeFieldResponse) SetStatusCode(v int32) *DescribeFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFieldResponse) SetBody(v *DescribeFieldResponseBody) *DescribeFieldResponse {
	s.Body = v
	return s
}

type DescribeGroupProductionsRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1182415068150980
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeGroupProductionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupProductionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsRequest) SetLang(v string) *DescribeGroupProductionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupProductionsRequest) SetRoleFor(v int64) *DescribeGroupProductionsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeGroupProductionsRequest) SetRoleType(v string) *DescribeGroupProductionsRequest {
	s.RoleType = &v
	return s
}

type DescribeGroupProductionsResponseBody struct {
	Data []*DescribeGroupProductionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Page *DescribeGroupProductionsResponseBodyPage   `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 358E012F-B516-599D-9ED0-A1A361CDE615
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupProductionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupProductionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsResponseBody) SetData(v []*DescribeGroupProductionsResponseBodyData) *DescribeGroupProductionsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeGroupProductionsResponseBody) SetPage(v *DescribeGroupProductionsResponseBodyPage) *DescribeGroupProductionsResponseBody {
	s.Page = v
	return s
}

func (s *DescribeGroupProductionsResponseBody) SetRequestId(v string) *DescribeGroupProductionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeGroupProductionsResponseBodyData struct {
	GroupName   *string                                                `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Productions []*DescribeGroupProductionsResponseBodyDataProductions `json:"Productions,omitempty" xml:"Productions,omitempty" type:"Repeated"`
}

func (s DescribeGroupProductionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupProductionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsResponseBodyData) SetGroupName(v string) *DescribeGroupProductionsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyData) SetProductions(v []*DescribeGroupProductionsResponseBodyDataProductions) *DescribeGroupProductionsResponseBodyData {
	s.Productions = v
	return s
}

type DescribeGroupProductionsResponseBodyDataProductions struct {
	// example:
	//
	// Rds
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// rds.aliyuncs.com
	DefaultDomain *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// example:
	//
	// 2014-08-15
	DefaultVersion *string                                                          `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Description    *string                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	FullDomains    []*string                                                        `json:"FullDomains,omitempty" xml:"FullDomains,omitempty" type:"Repeated"`
	Group          *string                                                          `json:"Group,omitempty" xml:"Group,omitempty"`
	Name           *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyList     []*DescribeGroupProductionsResponseBodyDataProductionsPolicyList `json:"PolicyList,omitempty" xml:"PolicyList,omitempty" type:"Repeated"`
	// example:
	//
	// rds
	RamCode *string `json:"RamCode,omitempty" xml:"RamCode,omitempty"`
	// example:
	//
	// RDS
	ShortName *string `json:"ShortName,omitempty" xml:"ShortName,omitempty"`
	// example:
	//
	// next
	Source   *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	Versions []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s DescribeGroupProductionsResponseBodyDataProductions) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupProductionsResponseBodyDataProductions) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetCode(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.Code = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetDefaultDomain(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.DefaultDomain = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetDefaultVersion(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.DefaultVersion = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetDescription(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.Description = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetFullDomains(v []*string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.FullDomains = v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetGroup(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.Group = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetName(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.Name = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetPolicyList(v []*DescribeGroupProductionsResponseBodyDataProductionsPolicyList) *DescribeGroupProductionsResponseBodyDataProductions {
	s.PolicyList = v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetRamCode(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.RamCode = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetShortName(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.ShortName = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetSource(v string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.Source = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductions) SetVersions(v []*string) *DescribeGroupProductionsResponseBodyDataProductions {
	s.Versions = v
	return s
}

type DescribeGroupProductionsResponseBodyDataProductionsPolicyList struct {
	// example:
	//
	// AliyunRAMReadOnlyAccess
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeGroupProductionsResponseBodyDataProductionsPolicyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupProductionsResponseBodyDataProductionsPolicyList) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsResponseBodyDataProductionsPolicyList) SetPolicyName(v string) *DescribeGroupProductionsResponseBodyDataProductionsPolicyList {
	s.PolicyName = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyDataProductionsPolicyList) SetType(v string) *DescribeGroupProductionsResponseBodyDataProductionsPolicyList {
	s.Type = &v
	return s
}

type DescribeGroupProductionsResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGroupProductionsResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupProductionsResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsResponseBodyPage) SetPageNumber(v int32) *DescribeGroupProductionsResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyPage) SetPageSize(v int32) *DescribeGroupProductionsResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupProductionsResponseBodyPage) SetTotalCount(v int32) *DescribeGroupProductionsResponseBodyPage {
	s.TotalCount = &v
	return s
}

type DescribeGroupProductionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupProductionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupProductionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupProductionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsResponse) SetHeaders(v map[string]*string) *DescribeGroupProductionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupProductionsResponse) SetStatusCode(v int32) *DescribeGroupProductionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupProductionsResponse) SetBody(v *DescribeGroupProductionsResponseBody) *DescribeGroupProductionsResponse {
	s.Body = v
	return s
}

type DescribeLatestRecordSchemaRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// c5c88b5e-97ca-435d-8c20-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeLatestRecordSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLatestRecordSchemaRequest) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaRequest) SetLang(v string) *DescribeLatestRecordSchemaRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLatestRecordSchemaRequest) SetPlaybookUuid(v string) *DescribeLatestRecordSchemaRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribeLatestRecordSchemaResponseBody struct {
	// The output structure information of the playbook.
	PlaybookNodeSchema *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema `json:"PlaybookNodeSchema,omitempty" xml:"PlaybookNodeSchema,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 10B92EE1-4597-593B-A131-7A17D25EF5C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLatestRecordSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLatestRecordSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaResponseBody) SetPlaybookNodeSchema(v *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) *DescribeLatestRecordSchemaResponseBody {
	s.PlaybookNodeSchema = v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBody) SetRequestId(v string) *DescribeLatestRecordSchemaResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema struct {
	// The structure information.
	NodeSchema []*DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema `json:"NodeSchema,omitempty" xml:"NodeSchema,omitempty" type:"Repeated"`
}

func (s DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema) SetNodeSchema(v []*DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchema {
	s.NodeSchema = v
	return s
}

type DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema struct {
	// The action name of the component.
	//
	// example:
	//
	// formatedata
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// The name of the component.
	//
	// example:
	//
	// DataFormat
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// DataFormat_1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The output fields.
	OutputFields []*string `json:"OutputFields,omitempty" xml:"OutputFields,omitempty" type:"Repeated"`
}

func (s DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) SetActionName(v string) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	s.ActionName = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) SetComponentName(v string) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	s.ComponentName = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) SetNodeName(v string) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	s.NodeName = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema) SetOutputFields(v []*string) *DescribeLatestRecordSchemaResponseBodyPlaybookNodeSchemaNodeSchema {
	s.OutputFields = v
	return s
}

type DescribeLatestRecordSchemaResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLatestRecordSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLatestRecordSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLatestRecordSchemaResponse) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaResponse) SetHeaders(v map[string]*string) *DescribeLatestRecordSchemaResponse {
	s.Headers = v
	return s
}

func (s *DescribeLatestRecordSchemaResponse) SetStatusCode(v int32) *DescribeLatestRecordSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLatestRecordSchemaResponse) SetBody(v *DescribeLatestRecordSchemaResponseBody) *DescribeLatestRecordSchemaResponse {
	s.Body = v
	return s
}

type DescribeNodeParamTagsRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// python3_2
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac343acc-1a61-4084-9a1c-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeNodeParamTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeParamTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeParamTagsRequest) SetLang(v string) *DescribeNodeParamTagsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNodeParamTagsRequest) SetNodeName(v string) *DescribeNodeParamTagsRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeNodeParamTagsRequest) SetPlaybookUuid(v string) *DescribeNodeParamTagsRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribeNodeParamTagsResponseBody struct {
	// The configuration of the recommended path.
	ParamReferredPaths []*DescribeNodeParamTagsResponseBodyParamReferredPaths `json:"ParamReferredPaths,omitempty" xml:"ParamReferredPaths,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6BE94351-712A-505D-A40A-BC77CC8254A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNodeParamTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeParamTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeParamTagsResponseBody) SetParamReferredPaths(v []*DescribeNodeParamTagsResponseBodyParamReferredPaths) *DescribeNodeParamTagsResponseBody {
	s.ParamReferredPaths = v
	return s
}

func (s *DescribeNodeParamTagsResponseBody) SetRequestId(v string) *DescribeNodeParamTagsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNodeParamTagsResponseBodyParamReferredPaths struct {
	// The name of the upstream node.
	//
	// example:
	//
	// DataFormat_1
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// The paths.
	ReferredPath []*string `json:"ReferredPath,omitempty" xml:"ReferredPath,omitempty" type:"Repeated"`
}

func (s DescribeNodeParamTagsResponseBodyParamReferredPaths) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeParamTagsResponseBodyParamReferredPaths) GoString() string {
	return s.String()
}

func (s *DescribeNodeParamTagsResponseBodyParamReferredPaths) SetParamName(v string) *DescribeNodeParamTagsResponseBodyParamReferredPaths {
	s.ParamName = &v
	return s
}

func (s *DescribeNodeParamTagsResponseBodyParamReferredPaths) SetReferredPath(v []*string) *DescribeNodeParamTagsResponseBodyParamReferredPaths {
	s.ReferredPath = v
	return s
}

type DescribeNodeParamTagsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeParamTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodeParamTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeParamTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeParamTagsResponse) SetHeaders(v map[string]*string) *DescribeNodeParamTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeParamTagsResponse) SetStatusCode(v int32) *DescribeNodeParamTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeParamTagsResponse) SetBody(v *DescribeNodeParamTagsResponseBody) *DescribeNodeParamTagsResponse {
	s.Body = v
	return s
}

type DescribeNodeUsedInfosRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The node name of the component.
	//
	// This parameter is required.
	//
	// example:
	//
	// python3_2
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac343acc-1a61-4084-9a1c-xxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeNodeUsedInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeUsedInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeUsedInfosRequest) SetLang(v string) *DescribeNodeUsedInfosRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNodeUsedInfosRequest) SetNodeName(v string) *DescribeNodeUsedInfosRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeNodeUsedInfosRequest) SetPlaybookUuid(v string) *DescribeNodeUsedInfosRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribeNodeUsedInfosResponseBody struct {
	// The node reference information. The value is in the JSON format and contains the following fields:
	//
	// 	- **action**: the referencing action. This field contains the following information:
	//
	//     	- **name**: the name of the referencing node.
	//
	//     	- **inputParams**: the parameter settings of the referencing node.
	//
	// example:
	//
	// {
	//
	//     "action": [
	//
	//         {
	//
	//             "name": "query_books",
	//
	//             "inputParams": [
	//
	//                 {
	//
	//                     "referInfos": [
	//
	//                         "${play_group.datalist.*.ids}"
	//
	//                     ],
	//
	//                     "name": "querySql"
	//
	//                 }
	//
	//             ]
	//
	//         }
	//
	//     ]
	//
	// }
	NodeUsedInfos *string `json:"NodeUsedInfos,omitempty" xml:"NodeUsedInfos,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3B10F836-C2B1-54FA-AB59-7591B548FB59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNodeUsedInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeUsedInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeUsedInfosResponseBody) SetNodeUsedInfos(v string) *DescribeNodeUsedInfosResponseBody {
	s.NodeUsedInfos = &v
	return s
}

func (s *DescribeNodeUsedInfosResponseBody) SetRequestId(v string) *DescribeNodeUsedInfosResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNodeUsedInfosResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeUsedInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodeUsedInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeUsedInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeUsedInfosResponse) SetHeaders(v map[string]*string) *DescribeNodeUsedInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeUsedInfosResponse) SetStatusCode(v int32) *DescribeNodeUsedInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeUsedInfosResponse) SetBody(v *DescribeNodeUsedInfosResponseBody) *DescribeNodeUsedInfosResponse {
	s.Body = v
	return s
}

type DescribeNotifyTemplateListRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 137602425xxx8726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeNotifyTemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotifyTemplateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNotifyTemplateListRequest) SetLang(v string) *DescribeNotifyTemplateListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNotifyTemplateListRequest) SetRoleFor(v int64) *DescribeNotifyTemplateListRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeNotifyTemplateListRequest) SetRoleType(v string) *DescribeNotifyTemplateListRequest {
	s.RoleType = &v
	return s
}

type DescribeNotifyTemplateListResponseBody struct {
	Data []*DescribeNotifyTemplateListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Page *DescribeNotifyTemplateListResponseBodyPage   `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// B3FED5B9-190A-5952-93A4-24FBF0F0C573
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNotifyTemplateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotifyTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNotifyTemplateListResponseBody) SetData(v []*DescribeNotifyTemplateListResponseBodyData) *DescribeNotifyTemplateListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNotifyTemplateListResponseBody) SetPage(v *DescribeNotifyTemplateListResponseBodyPage) *DescribeNotifyTemplateListResponseBody {
	s.Page = v
	return s
}

func (s *DescribeNotifyTemplateListResponseBody) SetRequestId(v string) *DescribeNotifyTemplateListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNotifyTemplateListResponseBodyData struct {
	// example:
	//
	// Dear $aliyunUID : Cloud Security Center Threat Analysis and Response has detected a newly discovered security incident $incidentName(Incident id :$incidentID) in $startTime, Please go to Cloud Security Center Console View.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// yundun_soar_incident_generate
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// [\\"aliyunUID\\",\\"incidentName\\",\\"incidentID\\",\\"startTime\\"]
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// [Alibaba Cloud Threat Analysis and Response] has detected a newly discovered security incident $incidentName($incidentID)
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// example:
	//
	// incident generate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeNotifyTemplateListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotifyTemplateListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNotifyTemplateListResponseBodyData) SetContent(v string) *DescribeNotifyTemplateListResponseBodyData {
	s.Content = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyData) SetEventId(v string) *DescribeNotifyTemplateListResponseBodyData {
	s.EventId = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyData) SetParams(v string) *DescribeNotifyTemplateListResponseBodyData {
	s.Params = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyData) SetSubject(v string) *DescribeNotifyTemplateListResponseBodyData {
	s.Subject = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyData) SetTemplateName(v string) *DescribeNotifyTemplateListResponseBodyData {
	s.TemplateName = &v
	return s
}

type DescribeNotifyTemplateListResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNotifyTemplateListResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotifyTemplateListResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeNotifyTemplateListResponseBodyPage) SetPageNumber(v int32) *DescribeNotifyTemplateListResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyPage) SetPageSize(v int32) *DescribeNotifyTemplateListResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeNotifyTemplateListResponseBodyPage) SetTotalCount(v int32) *DescribeNotifyTemplateListResponseBodyPage {
	s.TotalCount = &v
	return s
}

type DescribeNotifyTemplateListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNotifyTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNotifyTemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotifyTemplateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNotifyTemplateListResponse) SetHeaders(v map[string]*string) *DescribeNotifyTemplateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNotifyTemplateListResponse) SetStatusCode(v int32) *DescribeNotifyTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNotifyTemplateListResponse) SetBody(v *DescribeNotifyTemplateListResponseBody) *DescribeNotifyTemplateListResponse {
	s.Body = v
	return s
}

type DescribeOpenApiInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DescribePopApiItemList
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2018-12-03
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Sas
	PopCode *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
	// example:
	//
	// 1592757xxx002956
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeOpenApiInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiInfoRequest) SetApiName(v string) *DescribeOpenApiInfoRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeOpenApiInfoRequest) SetApiVersion(v string) *DescribeOpenApiInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribeOpenApiInfoRequest) SetLang(v string) *DescribeOpenApiInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOpenApiInfoRequest) SetPopCode(v string) *DescribeOpenApiInfoRequest {
	s.PopCode = &v
	return s
}

func (s *DescribeOpenApiInfoRequest) SetRoleFor(v int64) *DescribeOpenApiInfoRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeOpenApiInfoRequest) SetRoleType(v string) *DescribeOpenApiInfoRequest {
	s.RoleType = &v
	return s
}

type DescribeOpenApiInfoResponseBody struct {
	Data *DescribeOpenApiInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 358E012F-B516-599D-9ED0-A1A361CDE615
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenApiInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiInfoResponseBody) SetData(v *DescribeOpenApiInfoResponseBodyData) *DescribeOpenApiInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenApiInfoResponseBody) SetRequestId(v string) *DescribeOpenApiInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOpenApiInfoResponseBodyData struct {
	// example:
	//
	// describeEcs
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {}
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// example:
	//
	// []
	OutputParams *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	// example:
	//
	// []
	ResponseDemo *string `json:"ResponseDemo,omitempty" xml:"ResponseDemo,omitempty"`
	// example:
	//
	// describeEcs
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// describeEcs
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeOpenApiInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiInfoResponseBodyData) SetDescription(v string) *DescribeOpenApiInfoResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeOpenApiInfoResponseBodyData) SetInputParams(v string) *DescribeOpenApiInfoResponseBodyData {
	s.InputParams = &v
	return s
}

func (s *DescribeOpenApiInfoResponseBodyData) SetOutputParams(v string) *DescribeOpenApiInfoResponseBodyData {
	s.OutputParams = &v
	return s
}

func (s *DescribeOpenApiInfoResponseBodyData) SetResponseDemo(v string) *DescribeOpenApiInfoResponseBodyData {
	s.ResponseDemo = &v
	return s
}

func (s *DescribeOpenApiInfoResponseBodyData) SetSummary(v string) *DescribeOpenApiInfoResponseBodyData {
	s.Summary = &v
	return s
}

func (s *DescribeOpenApiInfoResponseBodyData) SetTitle(v string) *DescribeOpenApiInfoResponseBodyData {
	s.Title = &v
	return s
}

type DescribeOpenApiInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenApiInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenApiInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiInfoResponse) SetHeaders(v map[string]*string) *DescribeOpenApiInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenApiInfoResponse) SetStatusCode(v int32) *DescribeOpenApiInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenApiInfoResponse) SetBody(v *DescribeOpenApiInfoResponseBody) *DescribeOpenApiInfoResponse {
	s.Body = v
	return s
}

type DescribeOpenApiListRequest struct {
	// example:
	//
	// DescribePopApiItemList
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-10-01
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Sas
	PopCode *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
	// example:
	//
	// 137602xxx8718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeOpenApiListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiListRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiListRequest) SetApiName(v string) *DescribeOpenApiListRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeOpenApiListRequest) SetApiVersion(v string) *DescribeOpenApiListRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribeOpenApiListRequest) SetLang(v string) *DescribeOpenApiListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOpenApiListRequest) SetPopCode(v string) *DescribeOpenApiListRequest {
	s.PopCode = &v
	return s
}

func (s *DescribeOpenApiListRequest) SetRoleFor(v int64) *DescribeOpenApiListRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeOpenApiListRequest) SetRoleType(v string) *DescribeOpenApiListRequest {
	s.RoleType = &v
	return s
}

type DescribeOpenApiListResponseBody struct {
	Data *DescribeOpenApiListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// EF2ECA2D-D8E6-5021-BF5C-19DD6D52C5B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenApiListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiListResponseBody) SetData(v *DescribeOpenApiListResponseBodyData) *DescribeOpenApiListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenApiListResponseBody) SetRequestId(v string) *DescribeOpenApiListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOpenApiListResponseBodyData struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// [{"apis":[{"summary":"get account information","deprecated":false,"name":"DescAccountSummary","title":"get account information"}],"childrens":[],"title":"account"}]
	Directories interface{} `json:"Directories,omitempty" xml:"Directories,omitempty"`
	// example:
	//
	// 2018-12-03
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeOpenApiListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiListResponseBodyData) SetCode(v string) *DescribeOpenApiListResponseBodyData {
	s.Code = &v
	return s
}

func (s *DescribeOpenApiListResponseBodyData) SetDirectories(v interface{}) *DescribeOpenApiListResponseBodyData {
	s.Directories = v
	return s
}

func (s *DescribeOpenApiListResponseBodyData) SetVersion(v string) *DescribeOpenApiListResponseBodyData {
	s.Version = &v
	return s
}

type DescribeOpenApiListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenApiListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenApiListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiListResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiListResponse) SetHeaders(v map[string]*string) *DescribeOpenApiListResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenApiListResponse) SetStatusCode(v int32) *DescribeOpenApiListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenApiListResponse) SetBody(v *DescribeOpenApiListResponseBody) *DescribeOpenApiListResponse {
	s.Body = v
	return s
}

type DescribePlaybookRequest struct {
	// The flag that indicates whether the playbook is of the debugging or published version. Valid values:
	//
	// 	- **1**: playbook of the debugging version
	//
	// 	- **0**: playbook of the published version
	//
	// example:
	//
	// 0
	DebugFlag *int32 `json:"DebugFlag,omitempty" xml:"DebugFlag,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9030076b-6733-4842-b05a-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The MD5 hash value of the playbook.
	//
	// example:
	//
	// 7a8f608dc64c242632aa578xxxxx
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
}

func (s DescribePlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookRequest) SetDebugFlag(v int32) *DescribePlaybookRequest {
	s.DebugFlag = &v
	return s
}

func (s *DescribePlaybookRequest) SetLang(v string) *DescribePlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookRequest) SetPlaybookUuid(v string) *DescribePlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookRequest) SetTaskflowMd5(v string) *DescribePlaybookRequest {
	s.TaskflowMd5 = &v
	return s
}

type DescribePlaybookResponseBody struct {
	// The configuration of the playbook.
	Playbook *DescribePlaybookResponseBodyPlaybook `json:"Playbook,omitempty" xml:"Playbook,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2989BC59-E9F0-5C83-B453-B368857649C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookResponseBody) SetPlaybook(v *DescribePlaybookResponseBodyPlaybook) *DescribePlaybookResponseBody {
	s.Playbook = v
	return s
}

func (s *DescribePlaybookResponseBody) SetRequestId(v string) *DescribePlaybookResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybookResponseBodyPlaybook struct {
	// The description of the playbook.
	//
	// example:
	//
	// demo playbook
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the playbook.
	//
	// example:
	//
	// demo_test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The number of times that the playbook failed to be run.
	//
	// example:
	//
	// 1
	FailExeNum *int32 `json:"FailExeNum,omitempty" xml:"FailExeNum,omitempty"`
	// The creation time of the playbook. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1665288858000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time of the playbook. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1677482519000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The input parameter configuration of the playbook. The value is a JSON array.
	//
	// >  For more information, see [DescribePlaybookInputOutput](~~DescribePlaybookInputOutput~~).
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "typeName": "String",
	//
	//         "dataClass": "normal",
	//
	//         "dataType": "String",
	//
	//         "description": "period",
	//
	//         "example": "",
	//
	//         "name": "period",
	//
	//         "required": false
	//
	//     }
	//
	// ]
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// The time when the playbook was last run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1665288858000
	LastExeTime *int64 `json:"LastExeTime,omitempty" xml:"LastExeTime,omitempty"`
	// The status of the playbook. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 0
	OnlineActive *bool `json:"OnlineActive,omitempty" xml:"OnlineActive,omitempty"`
	// The MD5 hash value in the latest published version of the playbook.
	//
	// example:
	//
	// asdfsdfe232-e2b2-44fd-b2cc-xxxxx
	OnlineReleaseTaskflowMd5 *string `json:"OnlineReleaseTaskflowMd5,omitempty" xml:"OnlineReleaseTaskflowMd5,omitempty"`
	// The type of the playbook. Valid values:
	//
	// 	- **preset**: predefined playbook
	//
	// 	- **user**: custom playbook
	//
	// example:
	//
	// preset
	OwnType *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// 8db257d3-e2b2-44fd-b2cc-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The number of times that the playbook was successfully run.
	//
	// example:
	//
	// 100
	SuccessExeNum *int32 `json:"SuccessExeNum,omitempty" xml:"SuccessExeNum,omitempty"`
	// The XML configuration of the playbook.
	//
	// example:
	//
	// <?xml version="1.0" encoding="UTF-8"?><bpmn:definitions xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:bpmn="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:dc="http://www.omg.org/spec/DD/20100524/DC" targetNamespace="http://bpmn.io/schema/bpmn" id="Definitions_1"><bpmn:process id="Process_1" isExecutable="false"><bpmn:startEvent id="StartEvent_1"/></bpmn:process><bpmndi:BPMNDiagram id="BPMNDiagram_1"><bpmndi:BPMNPlane id="BPMNPlane_1" bpmnElement="Process_1"><bpmndi:BPMNShape id="_BPMNShape_StartEvent_2" bpmnElement="StartEvent_1"><dc:Bounds height="36.0" width="36.0" x="173.0" y="102.0"/></bpmndi:BPMNShape></bpmndi:BPMNPlane></bpmndi:BPMNDiagram></bpmn:definitions>
	Taskflow *string `json:"Taskflow,omitempty" xml:"Taskflow,omitempty"`
	// The playbook configuration type.
	//
	// 	- **xml**: XML format.
	//
	// 	- **x6**: JSON format.
	//
	// example:
	//
	// xml
	TaskflowType *string `json:"TaskflowType,omitempty" xml:"TaskflowType,omitempty"`
}

func (s DescribePlaybookResponseBodyPlaybook) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookResponseBodyPlaybook) GoString() string {
	return s.String()
}

func (s *DescribePlaybookResponseBodyPlaybook) SetDescription(v string) *DescribePlaybookResponseBodyPlaybook {
	s.Description = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetDisplayName(v string) *DescribePlaybookResponseBodyPlaybook {
	s.DisplayName = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetFailExeNum(v int32) *DescribePlaybookResponseBodyPlaybook {
	s.FailExeNum = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetGmtCreate(v string) *DescribePlaybookResponseBodyPlaybook {
	s.GmtCreate = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetGmtModified(v string) *DescribePlaybookResponseBodyPlaybook {
	s.GmtModified = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetInputParams(v string) *DescribePlaybookResponseBodyPlaybook {
	s.InputParams = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetLastExeTime(v int64) *DescribePlaybookResponseBodyPlaybook {
	s.LastExeTime = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetOnlineActive(v bool) *DescribePlaybookResponseBodyPlaybook {
	s.OnlineActive = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetOnlineReleaseTaskflowMd5(v string) *DescribePlaybookResponseBodyPlaybook {
	s.OnlineReleaseTaskflowMd5 = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetOwnType(v string) *DescribePlaybookResponseBodyPlaybook {
	s.OwnType = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetPlaybookUuid(v string) *DescribePlaybookResponseBodyPlaybook {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetSuccessExeNum(v int32) *DescribePlaybookResponseBodyPlaybook {
	s.SuccessExeNum = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetTaskflow(v string) *DescribePlaybookResponseBodyPlaybook {
	s.Taskflow = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetTaskflowType(v string) *DescribePlaybookResponseBodyPlaybook {
	s.TaskflowType = &v
	return s
}

type DescribePlaybookResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookResponse) SetHeaders(v map[string]*string) *DescribePlaybookResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookResponse) SetStatusCode(v int32) *DescribePlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookResponse) SetBody(v *DescribePlaybookResponseBody) *DescribePlaybookResponse {
	s.Body = v
	return s
}

type DescribePlaybookInputOutputRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// b724d2b0-3c3b-4223-9bfd-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookInputOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookInputOutputRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookInputOutputRequest) SetLang(v string) *DescribePlaybookInputOutputRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookInputOutputRequest) SetPlaybookUuid(v string) *DescribePlaybookInputOutputRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribePlaybookInputOutputResponseBody struct {
	// The configurations.
	Config *DescribePlaybookInputOutputResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 688B4CCD-5272-5DCF-9D76-FE5EFEF545F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookInputOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookInputOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookInputOutputResponseBody) SetConfig(v *DescribePlaybookInputOutputResponseBodyConfig) *DescribePlaybookInputOutputResponseBody {
	s.Config = v
	return s
}

func (s *DescribePlaybookInputOutputResponseBody) SetRequestId(v string) *DescribePlaybookInputOutputResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybookInputOutputResponseBodyConfig struct {
	// The execution method of the playbook is in JSONObject format.
	ExeConfig *string `json:"ExeConfig,omitempty" xml:"ExeConfig,omitempty"`
	// The input parameter configuration of the playbook. The value is a JSON array.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "typeName": "String",
	//
	//         "dataClass": "normal",
	//
	//         "dataType": "String",
	//
	//         "description": "period",
	//
	//         "example": "",
	//
	//         "name": "period",
	//
	//         "required": false
	//
	//     }
	//
	// ]
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// The output parameter configuration. This parameter is unavailable and is always left empty.
	//
	// example:
	//
	// []
	OutputParams *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	// The input parameter type of the playbook. Valid values:
	//
	// 	- **template-ip**
	//
	// 	- **template-file**
	//
	// 	- **template-process**
	//
	// 	- **custom**
	//
	// example:
	//
	// custom
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// 9030076b-6733-4842-b05a-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookInputOutputResponseBodyConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookInputOutputResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetExeConfig(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.ExeConfig = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetInputParams(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.InputParams = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetOutputParams(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.OutputParams = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetParamType(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.ParamType = &v
	return s
}

func (s *DescribePlaybookInputOutputResponseBodyConfig) SetPlaybookUuid(v string) *DescribePlaybookInputOutputResponseBodyConfig {
	s.PlaybookUuid = &v
	return s
}

type DescribePlaybookInputOutputResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybookInputOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybookInputOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookInputOutputResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookInputOutputResponse) SetHeaders(v map[string]*string) *DescribePlaybookInputOutputResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookInputOutputResponse) SetStatusCode(v int32) *DescribePlaybookInputOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookInputOutputResponse) SetBody(v *DescribePlaybookInputOutputResponseBody) *DescribePlaybookInputOutputResponse {
	s.Body = v
	return s
}

type DescribePlaybookMetricsRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2a687089-d4dd-47d4-9709-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookMetricsRequest) SetLang(v string) *DescribePlaybookMetricsRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookMetricsRequest) SetPlaybookUuid(v string) *DescribePlaybookMetricsRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribePlaybookMetricsResponseBody struct {
	// The details of the playbook.
	Metrics *DescribePlaybookMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 567D3D0B-2153-5860-BF9A-F9DEED55FB73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookMetricsResponseBody) SetMetrics(v *DescribePlaybookMetricsResponseBodyMetrics) *DescribePlaybookMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *DescribePlaybookMetricsResponseBody) SetRequestId(v string) *DescribePlaybookMetricsResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybookMetricsResponseBodyMetrics struct {
	// The status of the playbook. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	Active *int32 `json:"Active,omitempty" xml:"Active,omitempty"`
	// The description of the playbook.
	//
	// example:
	//
	// This is a playbook for waf processing
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the playbook.
	//
	// example:
	//
	// demo name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The number of the tasks that are created for the playbook and failed to run.
	//
	// example:
	//
	// 10
	FailNum *int32 `json:"FailNum,omitempty" xml:"FailNum,omitempty"`
	// The time when the playbook was created. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1655277397000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The number of historical versions of the playbook.
	//
	// example:
	//
	// 10
	HistoryMd5 *int32 `json:"HistoryMd5,omitempty" xml:"HistoryMd5,omitempty"`
	// The time when the playbook was last run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1683526277415
	LastRuntime *int64 `json:"LastRuntime,omitempty" xml:"LastRuntime,omitempty"`
	// The type of the playbook. Valid values:
	//
	// 	- **preset**: predefined playbook
	//
	// 	- **user**: custom playbook
	//
	// example:
	//
	// user
	OwnType *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// 0fbc9bdb-9ae3-4ef4-a709-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The number of the tasks that are created for the playbook and were successfully run.
	//
	// example:
	//
	// 100
	SuccNum *int32 `json:"SuccNum,omitempty" xml:"SuccNum,omitempty"`
}

func (s DescribePlaybookMetricsResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetActive(v int32) *DescribePlaybookMetricsResponseBodyMetrics {
	s.Active = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetDescription(v string) *DescribePlaybookMetricsResponseBodyMetrics {
	s.Description = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetDisplayName(v string) *DescribePlaybookMetricsResponseBodyMetrics {
	s.DisplayName = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetFailNum(v int32) *DescribePlaybookMetricsResponseBodyMetrics {
	s.FailNum = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetGmtCreate(v int64) *DescribePlaybookMetricsResponseBodyMetrics {
	s.GmtCreate = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetHistoryMd5(v int32) *DescribePlaybookMetricsResponseBodyMetrics {
	s.HistoryMd5 = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetLastRuntime(v int64) *DescribePlaybookMetricsResponseBodyMetrics {
	s.LastRuntime = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetOwnType(v string) *DescribePlaybookMetricsResponseBodyMetrics {
	s.OwnType = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetPlaybookUuid(v string) *DescribePlaybookMetricsResponseBodyMetrics {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetSuccNum(v int32) *DescribePlaybookMetricsResponseBodyMetrics {
	s.SuccNum = &v
	return s
}

type DescribePlaybookMetricsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybookMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybookMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookMetricsResponse) SetHeaders(v map[string]*string) *DescribePlaybookMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookMetricsResponse) SetStatusCode(v int32) *DescribePlaybookMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookMetricsResponse) SetBody(v *DescribePlaybookMetricsResponseBody) *DescribePlaybookMetricsResponse {
	s.Body = v
	return s
}

type DescribePlaybookNodesOutputRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the component node.
	//
	// This parameter is required.
	//
	// example:
	//
	// DataFormat_1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac343acc-1a61-4084-9a1c-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookNodesOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNodesOutputRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNodesOutputRequest) SetLang(v string) *DescribePlaybookNodesOutputRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookNodesOutputRequest) SetNodeName(v string) *DescribePlaybookNodesOutputRequest {
	s.NodeName = &v
	return s
}

func (s *DescribePlaybookNodesOutputRequest) SetPlaybookUuid(v string) *DescribePlaybookNodesOutputRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribePlaybookNodesOutputResponseBody struct {
	// The output data of the component node.
	PlaybookNodesOutput *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput `json:"PlaybookNodesOutput,omitempty" xml:"PlaybookNodesOutput,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A491170C-FE1F-520E-83D4-72ED205B72ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookNodesOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNodesOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNodesOutputResponseBody) SetPlaybookNodesOutput(v *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) *DescribePlaybookNodesOutputResponseBody {
	s.PlaybookNodesOutput = v
	return s
}

func (s *DescribePlaybookNodesOutputResponseBody) SetRequestId(v string) *DescribePlaybookNodesOutputResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput struct {
	// The name of the component node.
	//
	// example:
	//
	// DataFormat_1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The historical output data of the component node. The value is in the JSON string format. If no data is found, the parameter is left empty.
	//
	// example:
	//
	// {
	//
	//     "datalist": [
	//
	//         {
	//
	//             "score": "10",
	//
	//             "ip": "1.1.1.1"
	//
	//         }
	//
	//     ],
	//
	//     "total_data_successful": 1,
	//
	//     "filter_total_data": 1,
	//
	//     "total_data": 1,
	//
	//     "total_exe_successful": 1,
	//
	//     "total_exe": 1,
	//
	//     "total_data_with_dup": 1,
	//
	//     "filter_total_data_successful": 1,
	//
	//     "status": true
	//
	// }
	NodeOutput *string `json:"NodeOutput,omitempty" xml:"NodeOutput,omitempty"`
}

func (s DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) SetNodeName(v string) *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput {
	s.NodeName = &v
	return s
}

func (s *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput) SetNodeOutput(v string) *DescribePlaybookNodesOutputResponseBodyPlaybookNodesOutput {
	s.NodeOutput = &v
	return s
}

type DescribePlaybookNodesOutputResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybookNodesOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybookNodesOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNodesOutputResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNodesOutputResponse) SetHeaders(v map[string]*string) *DescribePlaybookNodesOutputResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookNodesOutputResponse) SetStatusCode(v int32) *DescribePlaybookNodesOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookNodesOutputResponse) SetBody(v *DescribePlaybookNodesOutputResponseBody) *DescribePlaybookNodesOutputResponse {
	s.Body = v
	return s
}

type DescribePlaybookNumberMetricsRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribePlaybookNumberMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNumberMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNumberMetricsRequest) SetLang(v string) *DescribePlaybookNumberMetricsRequest {
	s.Lang = &v
	return s
}

type DescribePlaybookNumberMetricsResponseBody struct {
	// The statistics.
	Metrics *DescribePlaybookNumberMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D4CC979E-3D5B-5A6A-BC87-C93C9E861C7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookNumberMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNumberMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNumberMetricsResponseBody) SetMetrics(v *DescribePlaybookNumberMetricsResponseBodyMetrics) *DescribePlaybookNumberMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *DescribePlaybookNumberMetricsResponseBody) SetRequestId(v string) *DescribePlaybookNumberMetricsResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybookNumberMetricsResponseBodyMetrics struct {
	// The number of enabled playbooks.
	//
	// example:
	//
	// 50
	StartUpNum *int32 `json:"StartUpNum,omitempty" xml:"StartUpNum,omitempty"`
	// The total number of playbooks.
	//
	// example:
	//
	// 100
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s DescribePlaybookNumberMetricsResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNumberMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNumberMetricsResponseBodyMetrics) SetStartUpNum(v int32) *DescribePlaybookNumberMetricsResponseBodyMetrics {
	s.StartUpNum = &v
	return s
}

func (s *DescribePlaybookNumberMetricsResponseBodyMetrics) SetTotalNum(v int32) *DescribePlaybookNumberMetricsResponseBodyMetrics {
	s.TotalNum = &v
	return s
}

type DescribePlaybookNumberMetricsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybookNumberMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybookNumberMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookNumberMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNumberMetricsResponse) SetHeaders(v map[string]*string) *DescribePlaybookNumberMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookNumberMetricsResponse) SetStatusCode(v int32) *DescribePlaybookNumberMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookNumberMetricsResponse) SetBody(v *DescribePlaybookNumberMetricsResponseBody) *DescribePlaybookNumberMetricsResponse {
	s.Body = v
	return s
}

type DescribePlaybookReleasesRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Default value: 1. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. If you do not specify the PageSize parameter, 10 entries are returned by default.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac343acc-1a61-4084-9a1c-xxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookReleasesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookReleasesRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesRequest) SetLang(v string) *DescribePlaybookReleasesRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookReleasesRequest) SetPageNumber(v int32) *DescribePlaybookReleasesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePlaybookReleasesRequest) SetPageSize(v int32) *DescribePlaybookReleasesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlaybookReleasesRequest) SetPlaybookUuid(v string) *DescribePlaybookReleasesRequest {
	s.PlaybookUuid = &v
	return s
}

type DescribePlaybookReleasesResponseBody struct {
	// The pagination information.
	Page *DescribePlaybookReleasesResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The information about the playbook version.
	Records []*DescribePlaybookReleasesResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3DFBE11C-6EB6-5166-92D6-3397796AFE1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookReleasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookReleasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesResponseBody) SetPage(v *DescribePlaybookReleasesResponseBodyPage) *DescribePlaybookReleasesResponseBody {
	s.Page = v
	return s
}

func (s *DescribePlaybookReleasesResponseBody) SetRecords(v []*DescribePlaybookReleasesResponseBodyRecords) *DescribePlaybookReleasesResponseBody {
	s.Records = v
	return s
}

func (s *DescribePlaybookReleasesResponseBody) SetRequestId(v string) *DescribePlaybookReleasesResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybookReleasesResponseBodyPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePlaybookReleasesResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookReleasesResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesResponseBodyPage) SetPageNumber(v int32) *DescribePlaybookReleasesResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyPage) SetPageSize(v int32) *DescribePlaybookReleasesResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyPage) SetTotalCount(v int32) *DescribePlaybookReleasesResponseBodyPage {
	s.TotalCount = &v
	return s
}

type DescribePlaybookReleasesResponseBodyRecords struct {
	// The ID of the Alibaba Cloud account that is used to publish the version.
	//
	// example:
	//
	// 145xxxx985
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The description of the layer version.
	//
	// example:
	//
	// This is a new version
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the version was created. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1655277397000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the version was modified. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1691460804000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The record ID.
	//
	// example:
	//
	// 80xxx
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The MD5 value configured for the published version of the playbook.
	//
	// example:
	//
	// be0a4ef084dd174abe47xxxxx
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
}

func (s DescribePlaybookReleasesResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookReleasesResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetCreator(v string) *DescribePlaybookReleasesResponseBodyRecords {
	s.Creator = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetDescription(v string) *DescribePlaybookReleasesResponseBodyRecords {
	s.Description = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetGmtCreate(v int64) *DescribePlaybookReleasesResponseBodyRecords {
	s.GmtCreate = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetGmtModified(v int64) *DescribePlaybookReleasesResponseBodyRecords {
	s.GmtModified = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetId(v int32) *DescribePlaybookReleasesResponseBodyRecords {
	s.Id = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetTaskflowMd5(v string) *DescribePlaybookReleasesResponseBodyRecords {
	s.TaskflowMd5 = &v
	return s
}

type DescribePlaybookReleasesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybookReleasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybookReleasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybookReleasesResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesResponse) SetHeaders(v map[string]*string) *DescribePlaybookReleasesResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybookReleasesResponse) SetStatusCode(v int32) *DescribePlaybookReleasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybookReleasesResponse) SetBody(v *DescribePlaybookReleasesResponseBody) *DescribePlaybookReleasesResponse {
	s.Body = v
	return s
}

type DescribePlaybooksRequest struct {
	// Activation status of the playbook. Values:
	//
	// - **1**: Indicates the playbook is activated.
	//
	// - **0**: Indicates the playbook is not activated.
	//
	// example:
	//
	// 1
	Active *int32 `json:"Active,omitempty" xml:"Active,omitempty"`
	// End time for the query, in 13-digit timestamp format.
	//
	// example:
	//
	// 1683858064361
	EndMillis *int64 `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	// Specifies the language type for the request and response, default is **zh**. Values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the playbook.
	//
	// example:
	//
	// demo_playbook
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sorting logic, with a default value of **desc**. Values:
	//
	// - **desc**: Descending order.
	//
	// - **asc**: Ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// Type of the playbook. Values:
	//
	// - **preset**: Predefined playbook.
	//
	// - **user**: Custom playbook.
	//
	// example:
	//
	// user
	OwnType *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// Sets the page number from which to start displaying the query results. The default value is 1, indicating the first page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Specifies the maximum number of items to display per page in a paginated query. The default number of items per page is 20. If the PageSize parameter is empty, it will return 10 items by default.
	//
	// > It is recommended that the PageSize value is not empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The trigger method for the playbook, with a default value of **query all**. Values:
	//
	// - **template-incident**: Security incident.
	//
	// - **template-ip**: IP entity.
	//
	// - **template-file**: File entity.
	//
	// - **template-process**: Process entity.
	//
	// - **template-alert**: Security alert.
	//
	// - **template-domain**: Domain entity.
	//
	// - **template-container**: Container entity.
	//
	// - **template-host**: Host entity.
	//
	// - **template-custom**: Custom.
	//
	// example:
	//
	// template-alert
	ParamTypes *string `json:"ParamTypes,omitempty" xml:"ParamTypes,omitempty"`
	// The UUID of the playbook.
	//
	// > You can use the UUID to query specific playbook information.
	//
	// > - Call the [DescribePlaybooks](~~DescribePlaybooks~~) API to obtain this parameter.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-1xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// UUID List of the playbook.
	//
	// Note You can use the UUID list to query specific playbook information.
	//
	// Call the DescribePlaybooks API to obtain this parameter.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-1xxxxxx,7745e6cff-319e-4ede-97bc-1xxxxxx
	PlaybookUuids *string `json:"PlaybookUuids,omitempty" xml:"PlaybookUuids,omitempty"`
	// The sorting basis, with a default value of **1**. Values:
	//
	// - **1**: Last modified time.
	//
	// - **2**: Most recent execution time.
	//
	// example:
	//
	// 1
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Start time for the query, in 13-digit timestamp format.
	//
	// example:
	//
	// 1683526277415
	StartMillis *int64 `json:"StartMillis,omitempty" xml:"StartMillis,omitempty"`
}

func (s DescribePlaybooksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybooksRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybooksRequest) SetActive(v int32) *DescribePlaybooksRequest {
	s.Active = &v
	return s
}

func (s *DescribePlaybooksRequest) SetEndMillis(v int64) *DescribePlaybooksRequest {
	s.EndMillis = &v
	return s
}

func (s *DescribePlaybooksRequest) SetLang(v string) *DescribePlaybooksRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybooksRequest) SetName(v string) *DescribePlaybooksRequest {
	s.Name = &v
	return s
}

func (s *DescribePlaybooksRequest) SetOrder(v string) *DescribePlaybooksRequest {
	s.Order = &v
	return s
}

func (s *DescribePlaybooksRequest) SetOwnType(v string) *DescribePlaybooksRequest {
	s.OwnType = &v
	return s
}

func (s *DescribePlaybooksRequest) SetPageNumber(v int64) *DescribePlaybooksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePlaybooksRequest) SetPageSize(v int32) *DescribePlaybooksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlaybooksRequest) SetParamTypes(v string) *DescribePlaybooksRequest {
	s.ParamTypes = &v
	return s
}

func (s *DescribePlaybooksRequest) SetPlaybookUuid(v string) *DescribePlaybooksRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybooksRequest) SetPlaybookUuids(v string) *DescribePlaybooksRequest {
	s.PlaybookUuids = &v
	return s
}

func (s *DescribePlaybooksRequest) SetSort(v string) *DescribePlaybooksRequest {
	s.Sort = &v
	return s
}

func (s *DescribePlaybooksRequest) SetStartMillis(v int64) *DescribePlaybooksRequest {
	s.StartMillis = &v
	return s
}

type DescribePlaybooksResponseBody struct {
	// Pagination query information.
	Page *DescribePlaybooksResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// List of playbooks.
	Playbooks []*DescribePlaybooksResponseBodyPlaybooks `json:"Playbooks,omitempty" xml:"Playbooks,omitempty" type:"Repeated"`
	// The ID of the current request, generated by Alibaba Cloud as a unique identifier for troubleshooting and problem localization.
	//
	// example:
	//
	// 138B5AB7-****-5814-87A3-E3E****F207E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybooksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybooksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybooksResponseBody) SetPage(v *DescribePlaybooksResponseBodyPage) *DescribePlaybooksResponseBody {
	s.Page = v
	return s
}

func (s *DescribePlaybooksResponseBody) SetPlaybooks(v []*DescribePlaybooksResponseBodyPlaybooks) *DescribePlaybooksResponseBody {
	s.Playbooks = v
	return s
}

func (s *DescribePlaybooksResponseBody) SetRequestId(v string) *DescribePlaybooksResponseBody {
	s.RequestId = &v
	return s
}

type DescribePlaybooksResponseBodyPage struct {
	// The page number in pagination queries.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of items per page in pagination queries.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of items found.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePlaybooksResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybooksResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribePlaybooksResponseBodyPage) SetPageNumber(v int32) *DescribePlaybooksResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPage) SetPageSize(v int32) *DescribePlaybooksResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPage) SetTotalCount(v int32) *DescribePlaybooksResponseBodyPage {
	s.TotalCount = &v
	return s
}

type DescribePlaybooksResponseBodyPlaybooks struct {
	// The status indicator of the playbook. Values:
	//
	// - **1**: Indicates the playbook is activated.
	//
	// - **0**: Indicates the playbook is deactivated.
	//
	// example:
	//
	// 1
	Active *int32 `json:"Active,omitempty" xml:"Active,omitempty"`
	// The display name of the playbook.
	//
	// example:
	//
	// demo_playbook
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The creation time of the playbook, in 13-digit timestamp format.
	//
	// example:
	//
	// 1683526277415
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time of the playbook.
	//
	// example:
	//
	// 1681396398000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The last execution time of the playbook, in 13-digit timestamp format.
	//
	// example:
	//
	// 1683526277415
	LastRuntime *int64 `json:"LastRuntime,omitempty" xml:"LastRuntime,omitempty"`
	// The type of the playbook. Values:
	//
	// - **preset**: Predefined playbook.
	//
	// - **user**: Custom playbook.
	//
	// example:
	//
	// user
	OwnType *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The trigger method for the playbook, with a default value of **query all**. Possible values are:
	//
	// - **template-incident**: Security incident.
	//
	// - **template-ip**: IP entity.
	//
	// - **template-file**: File entity.
	//
	// - **template-process**: Process entity.
	//
	// - **template-alert**: Security alert.
	//
	// - **template-domain**: Domain entity.
	//
	// - **template-container**: Container entity.
	//
	// - **template-host**: Host entity.
	//
	// example:
	//
	// template-alert
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// bb5a8640-a14f-44ef-8376-cxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybooksResponseBodyPlaybooks) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybooksResponseBodyPlaybooks) GoString() string {
	return s.String()
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetActive(v int32) *DescribePlaybooksResponseBodyPlaybooks {
	s.Active = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetDisplayName(v string) *DescribePlaybooksResponseBodyPlaybooks {
	s.DisplayName = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetGmtCreate(v int64) *DescribePlaybooksResponseBodyPlaybooks {
	s.GmtCreate = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetGmtModified(v string) *DescribePlaybooksResponseBodyPlaybooks {
	s.GmtModified = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetLastRuntime(v int64) *DescribePlaybooksResponseBodyPlaybooks {
	s.LastRuntime = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetOwnType(v string) *DescribePlaybooksResponseBodyPlaybooks {
	s.OwnType = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetParamType(v string) *DescribePlaybooksResponseBodyPlaybooks {
	s.ParamType = &v
	return s
}

func (s *DescribePlaybooksResponseBodyPlaybooks) SetPlaybookUuid(v string) *DescribePlaybooksResponseBodyPlaybooks {
	s.PlaybookUuid = &v
	return s
}

type DescribePlaybooksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlaybooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlaybooksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlaybooksResponse) GoString() string {
	return s.String()
}

func (s *DescribePlaybooksResponse) SetHeaders(v map[string]*string) *DescribePlaybooksResponse {
	s.Headers = v
	return s
}

func (s *DescribePlaybooksResponse) SetStatusCode(v int32) *DescribePlaybooksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlaybooksResponse) SetBody(v *DescribePlaybooksResponseBody) *DescribePlaybooksResponse {
	s.Body = v
	return s
}

type DescribePopApiRequest struct {
	// The operation name of the Alibaba Cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// DescribeInstanceInfo
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The version number of the API.
	//
	// >  You can call the [DescribePopApiVersionList](~~DescribePopApiVersionList~~) operation to query the version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-10-01
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The environment in which the API operation parameter is used. Set the value to online.
	//
	// This parameter is required.
	//
	// example:
	//
	// online
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The POP code of the Alibaba Cloud service.
	//
	// >  You can call the [DescribeApiList](~~DescribeApiList~~) operation to query the POP code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sas
	PopCode *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
}

func (s DescribePopApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiRequest) GoString() string {
	return s.String()
}

func (s *DescribePopApiRequest) SetApiName(v string) *DescribePopApiRequest {
	s.ApiName = &v
	return s
}

func (s *DescribePopApiRequest) SetApiVersion(v string) *DescribePopApiRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribePopApiRequest) SetEnv(v string) *DescribePopApiRequest {
	s.Env = &v
	return s
}

func (s *DescribePopApiRequest) SetPopCode(v string) *DescribePopApiRequest {
	s.PopCode = &v
	return s
}

type DescribePopApiResponseBody struct {
	// The name of the API.
	//
	// example:
	//
	// AddAssetCleanConfig
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The information about the API.
	OpenApiMetaList []*DescribePopApiResponseBodyOpenApiMetaList `json:"OpenApiMetaList,omitempty" xml:"OpenApiMetaList,omitempty" type:"Repeated"`
	// The POP code of the Alibaba Cloud service.
	//
	// example:
	//
	// Sas
	PopCode *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A01B0BA-CFC4-5813-9EB0-A5DA15FA95AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The version of the API.
	//
	// example:
	//
	// 2019-09-10
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribePopApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePopApiResponseBody) SetApiName(v string) *DescribePopApiResponseBody {
	s.ApiName = &v
	return s
}

func (s *DescribePopApiResponseBody) SetOpenApiMetaList(v []*DescribePopApiResponseBodyOpenApiMetaList) *DescribePopApiResponseBody {
	s.OpenApiMetaList = v
	return s
}

func (s *DescribePopApiResponseBody) SetPopCode(v string) *DescribePopApiResponseBody {
	s.PopCode = &v
	return s
}

func (s *DescribePopApiResponseBody) SetRequestId(v string) *DescribePopApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePopApiResponseBody) SetVersion(v string) *DescribePopApiResponseBody {
	s.Version = &v
	return s
}

type DescribePopApiResponseBodyOpenApiMetaList struct {
	// The parameter description.
	//
	// example:
	//
	// demo parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The example value.
	//
	// example:
	//
	// 12.xx.xx.xx
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// DescribePopApi
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// The data type of the parameter field. Valid values:
	//
	// 	- **string**
	//
	// 	- **boolean**
	//
	// 	- **integer**
	//
	// 	- **long**
	//
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePopApiResponseBodyOpenApiMetaList) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiResponseBodyOpenApiMetaList) GoString() string {
	return s.String()
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetDescription(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Description = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetExampleValue(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.ExampleValue = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetName(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Name = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetRequired(v bool) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Required = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetType(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Type = &v
	return s
}

type DescribePopApiResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePopApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePopApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiResponse) GoString() string {
	return s.String()
}

func (s *DescribePopApiResponse) SetHeaders(v map[string]*string) *DescribePopApiResponse {
	s.Headers = v
	return s
}

func (s *DescribePopApiResponse) SetStatusCode(v int32) *DescribePopApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePopApiResponse) SetBody(v *DescribePopApiResponseBody) *DescribePopApiResponse {
	s.Body = v
	return s
}

type DescribePopApiItemListRequest struct {
	// The API operation name of the Alibaba Cloud service. Fuzzy match is supported.
	//
	// example:
	//
	// DescribePopApiItemList
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The version number of the API.
	//
	// >  You can call the [DescribePopApiVersionList](~~DescribePopApiVersionList~~) operation to query the version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-12-03
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The environment in which the API operation parameters are used. Set the value to online.
	//
	// This parameter is required.
	//
	// example:
	//
	// online
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The POP code of the Alibaba Cloud service.
	//
	// >  You can call the [DescribeApiList](~~DescribeApiList~~) operation to query the POP code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sas
	PopCode *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
}

func (s DescribePopApiItemListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiItemListRequest) GoString() string {
	return s.String()
}

func (s *DescribePopApiItemListRequest) SetApiName(v string) *DescribePopApiItemListRequest {
	s.ApiName = &v
	return s
}

func (s *DescribePopApiItemListRequest) SetApiVersion(v string) *DescribePopApiItemListRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribePopApiItemListRequest) SetEnv(v string) *DescribePopApiItemListRequest {
	s.Env = &v
	return s
}

func (s *DescribePopApiItemListRequest) SetLang(v string) *DescribePopApiItemListRequest {
	s.Lang = &v
	return s
}

func (s *DescribePopApiItemListRequest) SetPopCode(v string) *DescribePopApiItemListRequest {
	s.PopCode = &v
	return s
}

type DescribePopApiItemListResponseBody struct {
	// The names of API operations.
	Names []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// The POP code of the Alibaba Cloud service.
	//
	// example:
	//
	// Sas
	PopCode *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6336D603-7028-52DE-AD88-E34AA5248355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The version number of the API for the Alibaba Cloud service.
	//
	// example:
	//
	// 2018-12-03
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribePopApiItemListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiItemListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePopApiItemListResponseBody) SetNames(v []*string) *DescribePopApiItemListResponseBody {
	s.Names = v
	return s
}

func (s *DescribePopApiItemListResponseBody) SetPopCode(v string) *DescribePopApiItemListResponseBody {
	s.PopCode = &v
	return s
}

func (s *DescribePopApiItemListResponseBody) SetRequestId(v string) *DescribePopApiItemListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePopApiItemListResponseBody) SetTotal(v int64) *DescribePopApiItemListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribePopApiItemListResponseBody) SetVersion(v string) *DescribePopApiItemListResponseBody {
	s.Version = &v
	return s
}

type DescribePopApiItemListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePopApiItemListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePopApiItemListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePopApiItemListResponse) GoString() string {
	return s.String()
}

func (s *DescribePopApiItemListResponse) SetHeaders(v map[string]*string) *DescribePopApiItemListResponse {
	s.Headers = v
	return s
}

func (s *DescribePopApiItemListResponse) SetStatusCode(v int32) *DescribePopApiItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePopApiItemListResponse) SetBody(v *DescribePopApiItemListResponseBody) *DescribePopApiItemListResponse {
	s.Body = v
	return s
}

type DescribeProcessStatisticsRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1709821xxxxx3093
	RoleFor *string `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeProcessStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessStatisticsRequest) SetLang(v string) *DescribeProcessStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeProcessStatisticsRequest) SetRoleFor(v string) *DescribeProcessStatisticsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeProcessStatisticsRequest) SetRoleType(v string) *DescribeProcessStatisticsRequest {
	s.RoleType = &v
	return s
}

type DescribeProcessStatisticsResponseBody struct {
	Metrics *DescribeProcessStatisticsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// example:
	//
	// 4CFC0F8A-D600-5FFF-A0DF-3121C4C1B90F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProcessStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessStatisticsResponseBody) SetMetrics(v *DescribeProcessStatisticsResponseBodyMetrics) *DescribeProcessStatisticsResponseBody {
	s.Metrics = v
	return s
}

func (s *DescribeProcessStatisticsResponseBody) SetRequestId(v string) *DescribeProcessStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProcessStatisticsResponseBodyMetrics struct {
	// example:
	//
	// 1
	BanFileNum *int32 `json:"BanFileNum,omitempty" xml:"BanFileNum,omitempty"`
	// example:
	//
	// 1
	BanIpNum *int32 `json:"BanIpNum,omitempty" xml:"BanIpNum,omitempty"`
	// example:
	//
	// 1
	BanProcessNum *int32 `json:"BanProcessNum,omitempty" xml:"BanProcessNum,omitempty"`
	// example:
	//
	// 1
	TaskNum *int32 `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
}

func (s DescribeProcessStatisticsResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessStatisticsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) SetBanFileNum(v int32) *DescribeProcessStatisticsResponseBodyMetrics {
	s.BanFileNum = &v
	return s
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) SetBanIpNum(v int32) *DescribeProcessStatisticsResponseBodyMetrics {
	s.BanIpNum = &v
	return s
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) SetBanProcessNum(v int32) *DescribeProcessStatisticsResponseBodyMetrics {
	s.BanProcessNum = &v
	return s
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) SetTaskNum(v int32) *DescribeProcessStatisticsResponseBodyMetrics {
	s.TaskNum = &v
	return s
}

type DescribeProcessStatisticsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProcessStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProcessStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeProcessStatisticsResponse) SetHeaders(v map[string]*string) *DescribeProcessStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeProcessStatisticsResponse) SetStatusCode(v int32) *DescribeProcessStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProcessStatisticsResponse) SetBody(v *DescribeProcessStatisticsResponseBody) *DescribeProcessStatisticsResponse {
	s.Body = v
	return s
}

type DescribeProcessTaskCountRequest struct {
	// Collection of entity UUIDs.
	//
	// This parameter is required.
	EntityUuidList []*string `json:"EntityUuidList,omitempty" xml:"EntityUuidList,omitempty" type:"Repeated"`
	// Language type for request and response messages. Values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// User ID for administrators to switch to other member\\"s perspective.
	//
	// example:
	//
	// 104739******259
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// View type.
	//
	// - **0**: Current Alibaba Cloud account view.
	//
	// - **1**: View for all accounts under the enterprise.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeProcessTaskCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTaskCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessTaskCountRequest) SetEntityUuidList(v []*string) *DescribeProcessTaskCountRequest {
	s.EntityUuidList = v
	return s
}

func (s *DescribeProcessTaskCountRequest) SetLang(v string) *DescribeProcessTaskCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeProcessTaskCountRequest) SetRoleFor(v int64) *DescribeProcessTaskCountRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeProcessTaskCountRequest) SetRoleType(v string) *DescribeProcessTaskCountRequest {
	s.RoleType = &v
	return s
}

type DescribeProcessTaskCountResponseBody struct {
	// Transmitted data.
	Data []*DescribeProcessTaskCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of this call request, which is a unique identifier generated by Alibaba Cloud for this request, used for troubleshooting and problem localization.
	//
	// example:
	//
	// e866cce0-****-41de-817e-****8d5e2650
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProcessTaskCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTaskCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessTaskCountResponseBody) SetData(v []*DescribeProcessTaskCountResponseBodyData) *DescribeProcessTaskCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeProcessTaskCountResponseBody) SetRequestId(v string) *DescribeProcessTaskCountResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProcessTaskCountResponseBodyData struct {
	// Count.
	//
	// example:
	//
	// 67
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Entity UUID.
	//
	// example:
	//
	// a680c9ae-****-4c39-****-0302****fc8e
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
}

func (s DescribeProcessTaskCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTaskCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeProcessTaskCountResponseBodyData) SetCount(v int64) *DescribeProcessTaskCountResponseBodyData {
	s.Count = &v
	return s
}

func (s *DescribeProcessTaskCountResponseBodyData) SetEntityUuid(v string) *DescribeProcessTaskCountResponseBodyData {
	s.EntityUuid = &v
	return s
}

type DescribeProcessTaskCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProcessTaskCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProcessTaskCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTaskCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeProcessTaskCountResponse) SetHeaders(v map[string]*string) *DescribeProcessTaskCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeProcessTaskCountResponse) SetStatusCode(v int32) *DescribeProcessTaskCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProcessTaskCountResponse) SetBody(v *DescribeProcessTaskCountResponseBody) *DescribeProcessTaskCountResponse {
	s.Body = v
	return s
}

type DescribeProcessTasksRequest struct {
	// The sort order. Valid values:
	//
	// 	- **desc*	- (default)
	//
	// 	- **asc**
	//
	// example:
	//
	// desc
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The name of the handling entity.
	//
	// example:
	//
	// 127.0.0.1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The type of the handling entity. Valid values:
	//
	// 	- **ip**
	//
	// 	- **file**
	//
	// 	- **process**
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	EventUuid  *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// The field that you use to sort the result.
	//
	// >  You can obtain the field from the response result.
	//
	// example:
	//
	// gmtCreate
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The page number. Default value: 1. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. If you do not specify the PageSize parameter, 10 entries are returned by default.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The handling entity, handling scenario, or handling parameter that is used for fuzzy match.
	//
	// example:
	//
	// 12.x.x.x
	ParamContent *string `json:"ParamContent,omitempty" xml:"ParamContent,omitempty"`
	// The end of the time range for a handling task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessActionEnd *int64 `json:"ProcessActionEnd,omitempty" xml:"ProcessActionEnd,omitempty"`
	// The beginning of the time range for a handling task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessActionStart *int64 `json:"ProcessActionStart,omitempty" xml:"ProcessActionStart,omitempty"`
	// The end of the time range for an unblocking task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessRemoveEnd *int64 `json:"ProcessRemoveEnd,omitempty" xml:"ProcessRemoveEnd,omitempty"`
	// The beginning of the time range for an unblocking task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessRemoveStart *int64 `json:"ProcessRemoveStart,omitempty" xml:"ProcessRemoveStart,omitempty"`
	// The UUID of the handling policy.
	//
	// >  You can call the [ListDisposeStrategy](https://help.aliyun.com/document_detail/2584440.html) operation to query the UUID of the handling policy.
	//
	// example:
	//
	// 92af3c79-1754-4646-9366-9ddbd1e45536_xxxx
	ProcessStrategyUuid *string `json:"ProcessStrategyUuid,omitempty" xml:"ProcessStrategyUuid,omitempty"`
	// The scenario code of the handling task.
	//
	// >  You can call the [DescribeEnumItems](~~DescribeEnumItems~~) operation to query the scenario code of the handling task. This parameter is available when you set **EnumType*	- to **process**.
	//
	// example:
	//
	// event_xxx_whole_process
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The ID of the Alibaba Cloud account that is specified in the handling task.
	//
	// example:
	//
	// 125xxxxx9870
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The triggering source of the handling task. The value is a string array. Valid values:
	//
	// 	- **system**: triggered when you manually handle an event
	//
	// 	- **custom**: triggered by an event based on an automatic response rule
	//
	// 	- **custom_alert**: triggered by an alert based on an automatic response rule
	//
	// 	- **soar-manual**: triggered when you use SOAR to manually run a playbook
	//
	// 	- **soar-mdr**: triggered by Managed Security Service
	//
	// example:
	//
	// ["system"]
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique identifier of the handling task.
	//
	// >  This parameter is used to query a specific task. You can obtain the value from the response result.
	//
	// example:
	//
	// 150xxxxxxxxx95066
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The status of the handling task. The value is a string. Valid values:
	//
	// 	- **11**: being handled
	//
	// 	- **21**: being blocked
	//
	// 	- **22**: being quarantined
	//
	// 	- **23**: completed
	//
	// 	- **24**: added to the whitelist
	//
	// 	- **20**: successful
	//
	// 	- **90**: failed
	//
	// 	- **91**: unblocking failed
	//
	// 	- **92**: restoring quarantined files failed
	//
	// example:
	//
	// ["11","21"]
	TaskStatus    *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TriggerSource *string `json:"TriggerSource,omitempty" xml:"TriggerSource,omitempty"`
	// The cloud service that is associated with the handling task. The value is a string. Valid values:
	//
	// 	- **WAF**: Web Application Firewall (WAF)
	//
	// 	- **CFW**: Cloud Firewall
	//
	// 	- **Aegis**: Security Center
	//
	// example:
	//
	// ["WAF"]
	YunCode *string `json:"YunCode,omitempty" xml:"YunCode,omitempty"`
}

func (s DescribeProcessTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksRequest) SetDirection(v string) *DescribeProcessTasksRequest {
	s.Direction = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetEntityName(v string) *DescribeProcessTasksRequest {
	s.EntityName = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetEntityType(v string) *DescribeProcessTasksRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetEntityUuid(v string) *DescribeProcessTasksRequest {
	s.EntityUuid = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetEventUuid(v string) *DescribeProcessTasksRequest {
	s.EventUuid = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetOrderField(v string) *DescribeProcessTasksRequest {
	s.OrderField = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetPageNumber(v int64) *DescribeProcessTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetPageSize(v int32) *DescribeProcessTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetParamContent(v string) *DescribeProcessTasksRequest {
	s.ParamContent = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessActionEnd(v int64) *DescribeProcessTasksRequest {
	s.ProcessActionEnd = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessActionStart(v int64) *DescribeProcessTasksRequest {
	s.ProcessActionStart = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessRemoveEnd(v int64) *DescribeProcessTasksRequest {
	s.ProcessRemoveEnd = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessRemoveStart(v int64) *DescribeProcessTasksRequest {
	s.ProcessRemoveStart = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetProcessStrategyUuid(v string) *DescribeProcessTasksRequest {
	s.ProcessStrategyUuid = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetSceneCode(v string) *DescribeProcessTasksRequest {
	s.SceneCode = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetScope(v string) *DescribeProcessTasksRequest {
	s.Scope = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetSource(v string) *DescribeProcessTasksRequest {
	s.Source = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetTaskId(v string) *DescribeProcessTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetTaskStatus(v string) *DescribeProcessTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetTriggerSource(v string) *DescribeProcessTasksRequest {
	s.TriggerSource = &v
	return s
}

func (s *DescribeProcessTasksRequest) SetYunCode(v string) *DescribeProcessTasksRequest {
	s.YunCode = &v
	return s
}

type DescribeProcessTasksResponseBody struct {
	// The pagination information.
	Page *DescribeProcessTasksResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The handling tasks.
	ProcessTasks []*DescribeProcessTasksResponseBodyProcessTasks `json:"ProcessTasks,omitempty" xml:"ProcessTasks,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E7698CFB-4E1C-5840-8EC9-691B86729E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProcessTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksResponseBody) SetPage(v *DescribeProcessTasksResponseBodyPage) *DescribeProcessTasksResponseBody {
	s.Page = v
	return s
}

func (s *DescribeProcessTasksResponseBody) SetProcessTasks(v []*DescribeProcessTasksResponseBodyProcessTasks) *DescribeProcessTasksResponseBody {
	s.ProcessTasks = v
	return s
}

func (s *DescribeProcessTasksResponseBody) SetRequestId(v string) *DescribeProcessTasksResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProcessTasksResponseBodyPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProcessTasksResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTasksResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksResponseBodyPage) SetPageNumber(v int32) *DescribeProcessTasksResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyPage) SetPageSize(v int32) *DescribeProcessTasksResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyPage) SetTotalCount(v int32) *DescribeProcessTasksResponseBodyPage {
	s.TotalCount = &v
	return s
}

type DescribeProcessTasksResponseBodyProcessTasks struct {
	// The ID of the Alibaba Cloud account that is used to submit the handling task.
	//
	// example:
	//
	// 123xxxx355
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The name of the handling entity.
	//
	// example:
	//
	// 1.1.1.x
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The type of the handling entity.
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// sts_openapi.Info.DefenseSceneNotSupported
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// ParamError : The parameters of your request are invalid
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The error tip returned if the call failed.
	//
	// example:
	//
	// Verify that the input parameters of the components are correct
	ErrTip    *string `json:"ErrTip,omitempty" xml:"ErrTip,omitempty"`
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// The creation time of the handling task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	GmtCreateMillis *int64 `json:"GmtCreateMillis,omitempty" xml:"GmtCreateMillis,omitempty"`
	// The modification time of the handling task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	GmtModifiedMillis *int64 `json:"GmtModifiedMillis,omitempty" xml:"GmtModifiedMillis,omitempty"`
	// The input parameter of the handling task.
	//
	// example:
	//
	// {"groupuuid":"c6a9b1df-f4ac-4078-bef4-99xxxxxx"}
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// The ID of the associated policy.
	//
	// example:
	//
	// 92af3c79-1754-4646-9366-9ddbd1e45536_xxxx
	ProcessStrategyUuid *string `json:"ProcessStrategyUuid,omitempty" xml:"ProcessStrategyUuid,omitempty"`
	// The delivery time of the handling task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	ProcessTime *int64 `json:"ProcessTime,omitempty" xml:"ProcessTime,omitempty"`
	// The unblocking time of the handling task. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1700031183572
	RemoveTime *int64  `json:"RemoveTime,omitempty" xml:"RemoveTime,omitempty"`
	ReqUuid    *string `json:"ReqUuid,omitempty" xml:"ReqUuid,omitempty"`
	// The scenario code of the handling task.
	//
	// example:
	//
	// event_xxx_whole_process
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The scenario name of the handling task.
	//
	// example:
	//
	// waf_whole_process
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The ID of the Alibaba Cloud account that is specified in the handling task.
	//
	// example:
	//
	// 123xxxxx234
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The submission source of the handling task.
	//
	// example:
	//
	// system
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique identifier of the handling task.
	//
	// example:
	//
	// 150xxxxxxxxx95066
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The status of the handling task.
	//
	// example:
	//
	// 11
	TaskStatus    *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TriggerSource *string `json:"TriggerSource,omitempty" xml:"TriggerSource,omitempty"`
	// The code of the cloud service that is associated with the handling task.
	//
	// example:
	//
	// WAF
	YunCode *string `json:"YunCode,omitempty" xml:"YunCode,omitempty"`
}

func (s DescribeProcessTasksResponseBodyProcessTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTasksResponseBodyProcessTasks) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetCreator(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.Creator = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetEntityName(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.EntityName = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetEntityType(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.EntityType = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetEntityUuid(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.EntityUuid = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetErrCode(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ErrCode = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetErrMsg(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ErrMsg = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetErrTip(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ErrTip = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetEventUuid(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.EventUuid = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetGmtCreateMillis(v int64) *DescribeProcessTasksResponseBodyProcessTasks {
	s.GmtCreateMillis = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetGmtModifiedMillis(v int64) *DescribeProcessTasksResponseBodyProcessTasks {
	s.GmtModifiedMillis = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetInputParams(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.InputParams = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetProcessStrategyUuid(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ProcessStrategyUuid = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetProcessTime(v int64) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ProcessTime = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetRemoveTime(v int64) *DescribeProcessTasksResponseBodyProcessTasks {
	s.RemoveTime = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetReqUuid(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.ReqUuid = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetSceneCode(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.SceneCode = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetSceneName(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.SceneName = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetScope(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.Scope = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetSource(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.Source = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetTaskId(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetTaskStatus(v int32) *DescribeProcessTasksResponseBodyProcessTasks {
	s.TaskStatus = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetTriggerSource(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.TriggerSource = &v
	return s
}

func (s *DescribeProcessTasksResponseBodyProcessTasks) SetYunCode(v string) *DescribeProcessTasksResponseBodyProcessTasks {
	s.YunCode = &v
	return s
}

type DescribeProcessTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProcessTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProcessTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProcessTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeProcessTasksResponse) SetHeaders(v map[string]*string) *DescribeProcessTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeProcessTasksResponse) SetStatusCode(v int32) *DescribeProcessTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProcessTasksResponse) SetBody(v *DescribeProcessTasksResponseBody) *DescribeProcessTasksResponse {
	s.Body = v
	return s
}

type DescribeSoarRecordActionOutputListRequest struct {
	// The UUID of the component action.
	//
	// >  You can call the [DescribeSoarTaskAndActions](~~DescribeSoarTaskAndActions~~) operation to query the UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2202c90d-fa93-4726-bc32-xxxxxx
	ActionUuid *string `json:"ActionUuid,omitempty" xml:"ActionUuid,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Default value: 1. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. If you leave this parameter empty, 10 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSoarRecordActionOutputListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordActionOutputListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordActionOutputListRequest) SetActionUuid(v string) *DescribeSoarRecordActionOutputListRequest {
	s.ActionUuid = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListRequest) SetLang(v string) *DescribeSoarRecordActionOutputListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListRequest) SetPageNumber(v int32) *DescribeSoarRecordActionOutputListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListRequest) SetPageSize(v int32) *DescribeSoarRecordActionOutputListRequest {
	s.PageSize = &v
	return s
}

type DescribeSoarRecordActionOutputListResponseBody struct {
	// The data that is returned when the component action is performed. The value is a JSON array.
	//
	// >  The format of the output data is determined by the component that is configured when the playbook is written.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "a": "a",
	//
	//         "taskname": "92af3c79-1754-4646-9366-9ddbd1e45536_xxxx",
	//
	//         "log_time": 1699868849000
	//
	//     }
	//
	// ]
	ActionOutputs *string `json:"ActionOutputs,omitempty" xml:"ActionOutputs,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6A2BF9CF-3E32-5E45-A79B-8F67E0A4FE90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSoarRecordActionOutputListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordActionOutputListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetActionOutputs(v string) *DescribeSoarRecordActionOutputListResponseBody {
	s.ActionOutputs = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetPageNumber(v int32) *DescribeSoarRecordActionOutputListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetPageSize(v int32) *DescribeSoarRecordActionOutputListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetRequestId(v string) *DescribeSoarRecordActionOutputListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponseBody) SetTotalCount(v int32) *DescribeSoarRecordActionOutputListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSoarRecordActionOutputListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarRecordActionOutputListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarRecordActionOutputListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordActionOutputListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordActionOutputListResponse) SetHeaders(v map[string]*string) *DescribeSoarRecordActionOutputListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponse) SetStatusCode(v int32) *DescribeSoarRecordActionOutputListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarRecordActionOutputListResponse) SetBody(v *DescribeSoarRecordActionOutputListResponseBody) *DescribeSoarRecordActionOutputListResponse {
	s.Body = v
	return s
}

type DescribeSoarRecordInOutputRequest struct {
	// The UUID of the component action.
	//
	// >  You can call the [DescribeSoarTaskAndActions](~~DescribeSoarTaskAndActions~~) operation to query the UUIDs of component actions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0531ff66-dd05-4f24-84bf-xxxxxxxx
	ActionUuid *string `json:"ActionUuid,omitempty" xml:"ActionUuid,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeSoarRecordInOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordInOutputRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordInOutputRequest) SetActionUuid(v string) *DescribeSoarRecordInOutputRequest {
	s.ActionUuid = &v
	return s
}

func (s *DescribeSoarRecordInOutputRequest) SetLang(v string) *DescribeSoarRecordInOutputRequest {
	s.Lang = &v
	return s
}

type DescribeSoarRecordInOutputResponseBody struct {
	// The execution result of the component action.
	//
	// example:
	//
	// {
	//
	//     "actionUuid": "3896a25d-4967-493c-942e-4e60f27da1f7-xxxxx",
	//
	//     "outputSummary": {
	//
	//         "datalist": [
	//
	//             {
	//
	//                 "a": "a"
	//
	//             }
	//
	//         ],
	//
	//         "total_data_successful": 1,
	//
	//         "total_data": 1,
	//
	//         "total_exe_successful": 1,
	//
	//         "total_exe": 1,
	//
	//         "total_data_with_dup": 1,
	//
	//         "status": true
	//
	//     },
	//
	//     "outputSchema": {
	//
	//         "a": "String",
	//
	//         "startTime": "DateTime"
	//
	//     },
	//
	//     "inputParams": {
	//
	//         "inputData": [
	//
	//             {
	//
	//                 "outputFields": {
	//
	//                     "a": "a"
	//
	//                 }
	//
	//             }
	//
	//         ],
	//
	//         "totalSize": 1
	//
	//     },
	//
	//     "startTime": "2023-11-13 17:47:28.645",
	//
	//     "taskName": "92af3c79-1754-4646-9366-9ddbxxxxx"
	//
	// }
	InOutputInfo *string `json:"InOutputInfo,omitempty" xml:"InOutputInfo,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 372D8B41-AF8D-573A-9B3F-0924950F241F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSoarRecordInOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordInOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordInOutputResponseBody) SetInOutputInfo(v string) *DescribeSoarRecordInOutputResponseBody {
	s.InOutputInfo = &v
	return s
}

func (s *DescribeSoarRecordInOutputResponseBody) SetRequestId(v string) *DescribeSoarRecordInOutputResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSoarRecordInOutputResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarRecordInOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarRecordInOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordInOutputResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordInOutputResponse) SetHeaders(v map[string]*string) *DescribeSoarRecordInOutputResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarRecordInOutputResponse) SetStatusCode(v int32) *DescribeSoarRecordInOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarRecordInOutputResponse) SetBody(v *DescribeSoarRecordInOutputResponseBody) *DescribeSoarRecordInOutputResponse {
	s.Body = v
	return s
}

type DescribeSoarRecordsRequest struct {
	// The end time of the task execution, in 13-digit timestamp format.
	//
	// example:
	//
	// 1683772744953
	EndMillis *int64 `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	// Set the language type for requests and received messages. The default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Set which page to start displaying the query results from. The default value is 1, indicating the first page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Specify the maximum number of data entries per page when performing a paginated query. The default number of entries per page is 20. If the PageSize parameter is empty, it will return 10 entries by default.
	//
	// > It is recommended not to leave the PageSize value empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UUID of the playbook.
	//
	// > You can obtain this parameter by calling the [DescribePlaybooks](~~DescribePlaybooks~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8f55e76d-b5d5-4720-9cd7-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// UUID of the playbook task execution.
	//
	// > You can obtain this parameter by calling the [DescribeSoarRecords](https://help.aliyun.com/document_detail/2627455.html) interface.
	//
	// example:
	//
	// 6d412cfa-0905-4567-8a83-xxxxxx
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
	// The start time of the task execution, in 13-digit timestamp format.
	//
	// example:
	//
	// 1683526284584
	StartMillis *int64 `json:"StartMillis,omitempty" xml:"StartMillis,omitempty"`
	// The status of the task execution. Values:
	//
	// - **success**: Successful task.
	//
	// - **failed**: Failed task.
	//
	// - **inprogress**: Task in progress
	//
	// example:
	//
	// inprogress
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The MD5 value of the playbook configuration.
	//
	// example:
	//
	// be0a4ef084dd174abe478df52xxxxx
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
	// The Alibaba Cloud account ID that executed the playbook task.
	//
	// example:
	//
	// 127xxxx4392
	TriggerUser *string `json:"TriggerUser,omitempty" xml:"TriggerUser,omitempty"`
}

func (s DescribeSoarRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordsRequest) SetEndMillis(v int64) *DescribeSoarRecordsRequest {
	s.EndMillis = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetLang(v string) *DescribeSoarRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetPageNumber(v int32) *DescribeSoarRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetPageSize(v int32) *DescribeSoarRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetPlaybookUuid(v string) *DescribeSoarRecordsRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetRequestUuid(v string) *DescribeSoarRecordsRequest {
	s.RequestUuid = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetStartMillis(v int64) *DescribeSoarRecordsRequest {
	s.StartMillis = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetTaskStatus(v string) *DescribeSoarRecordsRequest {
	s.TaskStatus = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetTaskflowMd5(v string) *DescribeSoarRecordsRequest {
	s.TaskflowMd5 = &v
	return s
}

func (s *DescribeSoarRecordsRequest) SetTriggerUser(v string) *DescribeSoarRecordsRequest {
	s.TriggerUser = &v
	return s
}

type DescribeSoarRecordsResponseBody struct {
	// Information displayed on the page.
	//
	// This parameter is required.
	Page *DescribeSoarRecordsResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The ID of the current request, generated by Alibaba Cloud as a unique identifier for the request, which can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// 601C2DAC-6A67-5237-BEE8-5BF1CEE96296
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Execution record result set.
	SoarExecuteRecords []*DescribeSoarRecordsResponseBodySoarExecuteRecords `json:"SoarExecuteRecords,omitempty" xml:"SoarExecuteRecords,omitempty" type:"Repeated"`
}

func (s DescribeSoarRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordsResponseBody) SetPage(v *DescribeSoarRecordsResponseBodyPage) *DescribeSoarRecordsResponseBody {
	s.Page = v
	return s
}

func (s *DescribeSoarRecordsResponseBody) SetRequestId(v string) *DescribeSoarRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarRecordsResponseBody) SetSoarExecuteRecords(v []*DescribeSoarRecordsResponseBodySoarExecuteRecords) *DescribeSoarRecordsResponseBody {
	s.SoarExecuteRecords = v
	return s
}

type DescribeSoarRecordsResponseBodyPage struct {
	// The current page number in paginated queries.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of items per page in paginated queries.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of queried items.
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSoarRecordsResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordsResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordsResponseBodyPage) SetPageNumber(v int32) *DescribeSoarRecordsResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodyPage) SetPageSize(v int32) *DescribeSoarRecordsResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodyPage) SetTotalCount(v int32) *DescribeSoarRecordsResponseBodyPage {
	s.TotalCount = &v
	return s
}

type DescribeSoarRecordsResponseBodySoarExecuteRecords struct {
	// The end time of the component execution, in 13-digit timestamp format.
	//
	// example:
	//
	// 1686294686000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error message of the playbook task. This field is empty when the task succeeds.
	//
	// example:
	//
	// stime not match
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The request parameters of the playbook task.
	//
	// example:
	//
	// {
	//
	//     "input1": "xx.xx.xx.xx",
	//
	//     "input2": "7d"
	//
	// }
	RawEventReq *string `json:"RawEventReq,omitempty" xml:"RawEventReq,omitempty"`
	// The request ID of the playbook task, a unique ID for each task run.
	//
	// example:
	//
	// ba1ec480-aa90-4bb6-a1a7-9e311ae79321
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
	// The return information of the playbook, defined by the user within the playbook.
	//
	// example:
	//
	// Playbook finish
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// The start time of the task execution, in 13-digit timestamp format.
	//
	// example:
	//
	// 1675823338433
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the playbook task. Values:
	//
	// - **success**: Indicates successful execution.
	//
	// - **fail**: Indicates failed execution.
	//
	// - **running**: Indicates the task is running
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the playbook task, which is the same as the playbook\\"s UUID.
	//
	// example:
	//
	// 82848ebc-eaff-4791-acd4-xxxxx
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the playbook task, with values:
	//
	// - **general**: Represents a general playbook task.
	//
	// - **standard**: Represents a component execution task.
	//
	// example:
	//
	// standard
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The MD5 value of the playbook configuration.
	//
	// example:
	//
	// dea65a3db87fb9bd84bbxxxxx
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
	// The type of the playbook task. Values:
	//
	// - **debug**: Indicates a debugging task.
	//
	// - **manual**: Indicates a manual task.
	//
	// - **siem**: Indicates a task triggered by an event or alert.
	//
	// example:
	//
	// debug
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The Alibaba Cloud account ID that executes the playbook task.
	//
	// example:
	//
	// 127xxxx4392
	TriggerUser *string `json:"TriggerUser,omitempty" xml:"TriggerUser,omitempty"`
}

func (s DescribeSoarRecordsResponseBodySoarExecuteRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordsResponseBodySoarExecuteRecords) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetEndTime(v int64) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.EndTime = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetErrorMsg(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetRawEventReq(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.RawEventReq = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetRequestUuid(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.RequestUuid = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetResultMessage(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.ResultMessage = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetStartTime(v int64) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.StartTime = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetStatus(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.Status = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetTaskName(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.TaskName = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetTaskType(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.TaskType = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetTaskflowMd5(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.TaskflowMd5 = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetTriggerType(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.TriggerType = &v
	return s
}

func (s *DescribeSoarRecordsResponseBodySoarExecuteRecords) SetTriggerUser(v string) *DescribeSoarRecordsResponseBodySoarExecuteRecords {
	s.TriggerUser = &v
	return s
}

type DescribeSoarRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordsResponse) SetHeaders(v map[string]*string) *DescribeSoarRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarRecordsResponse) SetStatusCode(v int32) *DescribeSoarRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarRecordsResponse) SetBody(v *DescribeSoarRecordsResponseBody) *DescribeSoarRecordsResponse {
	s.Body = v
	return s
}

type DescribeSoarTaskAndActionsRequest struct {
	// The language of the content within the request and response.
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The playbook UUID.
	//
	// example:
	//
	// 1077f2f9-25e8-42d9-bfdf-1528e1313f6d
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
}

func (s DescribeSoarTaskAndActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarTaskAndActionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsRequest) SetLang(v string) *DescribeSoarTaskAndActionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSoarTaskAndActionsRequest) SetRequestUuid(v string) *DescribeSoarTaskAndActionsRequest {
	s.RequestUuid = &v
	return s
}

type DescribeSoarTaskAndActionsResponseBody struct {
	// The execution details of each task.
	Details *DescribeSoarTaskAndActionsResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 18017A93-3D5D-503A-8308-914543F1CBA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSoarTaskAndActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponseBody) SetDetails(v *DescribeSoarTaskAndActionsResponseBodyDetails) *DescribeSoarTaskAndActionsResponseBody {
	s.Details = v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBody) SetRequestId(v string) *DescribeSoarTaskAndActionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSoarTaskAndActionsResponseBodyDetails struct {
	// The list of component actions during the running of the playbook.
	Actions []*DescribeSoarTaskAndActionsResponseBodyDetailsActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The end of the time range during which the playbook is run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1699868848767
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error message of the task. If the task is successful, this field is empty.
	//
	// example:
	//
	// stime not match
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The request parameters of the task.
	//
	// example:
	//
	// {
	//
	//     "input1": "xx.xx.xx.xx",
	//
	//     "input2": "7d"
	//
	// }
	RawEventReq *string `json:"RawEventReq,omitempty" xml:"RawEventReq,omitempty"`
	// The request ID of the task. The value is unique.
	//
	// example:
	//
	// 17f75844-75cc-4174-86da-cec07a690142
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
	// The flag of the task. For debugging tasks, the value is **DEBUG**. For other tasks, the parameter is left empty.
	//
	// example:
	//
	// DEBUG
	ResultLevel *string `json:"ResultLevel,omitempty" xml:"ResultLevel,omitempty"`
	// The returned information about the playbook. You can define the value in the playbook.
	//
	// example:
	//
	// deubug playbook finished
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// The beginning of the time range during which the playbook is run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1699868848645
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// 	- **running**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The MD5 value of the playbook.
	//
	// example:
	//
	// ed127287-6699-4e4d-b986-9f770879xxx
	TaskFlowMd5 *string `json:"TaskFlowMd5,omitempty" xml:"TaskFlowMd5,omitempty"`
	// The name of the task. The value is the same as the playbook UUID.
	//
	// example:
	//
	// 92af3c79-1754-4646-9366-9ddbd1e45536
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The ID of the Alibaba Cloud account to which the task belongs.
	//
	// example:
	//
	// 127xxxx4392
	TaskTenantId *string `json:"TaskTenantId,omitempty" xml:"TaskTenantId,omitempty"`
	// The task type. Valid values:
	//
	// 	- **debug**: a debugging task
	//
	// 	- **manual**: a manual task
	//
	// 	- **siem**: an event-triggered task
	//
	// example:
	//
	// siem
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The ID of the Alibaba Cloud account that triggers the task.
	//
	// example:
	//
	// 127xxxx4392
	TriggerUser *string `json:"TriggerUser,omitempty" xml:"TriggerUser,omitempty"`
}

func (s DescribeSoarTaskAndActionsResponseBodyDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetActions(v []*DescribeSoarTaskAndActionsResponseBodyDetailsActions) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.Actions = v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetEndTime(v int64) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.EndTime = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetErrorMsg(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetRawEventReq(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.RawEventReq = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetRequestUuid(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.RequestUuid = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetResultLevel(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.ResultLevel = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetResultMessage(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.ResultMessage = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetStartTime(v int64) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.StartTime = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetStatus(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.Status = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTaskFlowMd5(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TaskFlowMd5 = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTaskName(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TaskName = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTaskTenantId(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TaskTenantId = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTriggerType(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TriggerType = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetails) SetTriggerUser(v string) *DescribeSoarTaskAndActionsResponseBodyDetails {
	s.TriggerUser = &v
	return s
}

type DescribeSoarTaskAndActionsResponseBodyDetailsActions struct {
	// The action name of the component.
	//
	// example:
	//
	// formatdata
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The UUID of the component execution record.
	//
	// example:
	//
	// 091be399-a937-4276-af78-xxxxxxxx
	ActionUuid *string `json:"ActionUuid,omitempty" xml:"ActionUuid,omitempty"`
	// The name of the asset that is used by the component.
	//
	// example:
	//
	// SLS Asset
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The component name.
	//
	// example:
	//
	// DataFormat
	Component *string `json:"Component,omitempty" xml:"Component,omitempty"`
	// The end of the time range during which the component is run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1699868848766
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The custom name of the node in the component.
	//
	// example:
	//
	// DataFormat_1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The request ID of the task. The value is unique.
	//
	// example:
	//
	// 8dac16c6-7411-4116-8d70-xxxxxxx
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
	// The beginning of the time range during which the component is run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1699868848731
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The running result of the component. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the task. The value is the same as the playbook UUID.
	//
	// example:
	//
	// ed127287-6699-4e4d-b986-xxxxxxx
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The status of the triggered component action.
	//
	// >  This parameter is disabled and left empty.
	//
	// example:
	//
	// NULL
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The ID of the Alibaba Cloud account that is used to execute the task.
	//
	// example:
	//
	// 127xxxx4392
	TriggerUser *string `json:"TriggerUser,omitempty" xml:"TriggerUser,omitempty"`
}

func (s DescribeSoarTaskAndActionsResponseBodyDetailsActions) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponseBodyDetailsActions) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetAction(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.Action = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetActionUuid(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.ActionUuid = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetAssetName(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.AssetName = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetComponent(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.Component = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetEndTime(v int64) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.EndTime = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetNodeName(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.NodeName = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetRequestUuid(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.RequestUuid = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetStartTime(v int64) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.StartTime = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetStatus(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.Status = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetTaskName(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.TaskName = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetTaskStatus(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.TaskStatus = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponseBodyDetailsActions) SetTriggerUser(v string) *DescribeSoarTaskAndActionsResponseBodyDetailsActions {
	s.TriggerUser = &v
	return s
}

type DescribeSoarTaskAndActionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarTaskAndActionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarTaskAndActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSoarTaskAndActionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsResponse) SetHeaders(v map[string]*string) *DescribeSoarTaskAndActionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarTaskAndActionsResponse) SetStatusCode(v int32) *DescribeSoarTaskAndActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarTaskAndActionsResponse) SetBody(v *DescribeSoarTaskAndActionsResponseBody) *DescribeSoarTaskAndActionsResponse {
	s.Body = v
	return s
}

type DescribeSophonCommandsRequest struct {
	// The name of the command. Fuzzy match is supported.
	//
	// example:
	//
	// waf_process
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeSophonCommandsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSophonCommandsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsRequest) SetName(v string) *DescribeSophonCommandsRequest {
	s.Name = &v
	return s
}

type DescribeSophonCommandsResponseBody struct {
	// The commands.
	Data []*DescribeSophonCommandsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1E1EC464-3BD7-518F-9937-BCC12E6855FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSophonCommandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSophonCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsResponseBody) SetData(v []*DescribeSophonCommandsResponseBodyData) *DescribeSophonCommandsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSophonCommandsResponseBody) SetRequestId(v string) *DescribeSophonCommandsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSophonCommandsResponseBodyData struct {
	// The description of the command.
	//
	// example:
	//
	// This is a action of processing for WAF
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the command.
	//
	// example:
	//
	// WAF Process IP
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The name of the command.
	//
	// example:
	//
	// waf_process_ip_v2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter configurations.
	ParamConfig []*DescribeSophonCommandsResponseBodyDataParamConfig `json:"ParamConfig,omitempty" xml:"ParamConfig,omitempty" type:"Repeated"`
}

func (s DescribeSophonCommandsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSophonCommandsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsResponseBodyData) SetDescription(v string) *DescribeSophonCommandsResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyData) SetDisplayName(v string) *DescribeSophonCommandsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyData) SetName(v string) *DescribeSophonCommandsResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyData) SetParamConfig(v []*DescribeSophonCommandsResponseBodyDataParamConfig) *DescribeSophonCommandsResponseBodyData {
	s.ParamConfig = v
	return s
}

type DescribeSophonCommandsResponseBodyDataParamConfig struct {
	// The regular expression that is used to check the format of the parameter value. If the parameter is left empty, the check is not performed.
	//
	// example:
	//
	// [0-9]{4}\\.[0-9]{4}\\.[0-9]{4}\\.[0-9]{4}
	CheckField *string `json:"CheckField,omitempty" xml:"CheckField,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// ip
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// Indicates whether the parameter is required. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Necessary *bool `json:"Necessary,omitempty" xml:"Necessary,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 12.xx.xx.xx
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSophonCommandsResponseBodyDataParamConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeSophonCommandsResponseBodyDataParamConfig) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) SetCheckField(v string) *DescribeSophonCommandsResponseBodyDataParamConfig {
	s.CheckField = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) SetField(v string) *DescribeSophonCommandsResponseBodyDataParamConfig {
	s.Field = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) SetNecessary(v bool) *DescribeSophonCommandsResponseBodyDataParamConfig {
	s.Necessary = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) SetValue(v string) *DescribeSophonCommandsResponseBodyDataParamConfig {
	s.Value = &v
	return s
}

type DescribeSophonCommandsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSophonCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSophonCommandsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSophonCommandsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsResponse) SetHeaders(v map[string]*string) *DescribeSophonCommandsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSophonCommandsResponse) SetStatusCode(v int32) *DescribeSophonCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSophonCommandsResponse) SetBody(v *DescribeSophonCommandsResponseBody) *DescribeSophonCommandsResponse {
	s.Body = v
	return s
}

type DescribeVendorApiListRequest struct {
	// example:
	//
	// AddAssetCleanConfig
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// example:
	//
	// Create
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// waf
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// Azure
	VendorCode *string `json:"VendorCode,omitempty" xml:"VendorCode,omitempty"`
}

func (s DescribeVendorApiListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVendorApiListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVendorApiListRequest) SetApiName(v string) *DescribeVendorApiListRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeVendorApiListRequest) SetKeyWord(v string) *DescribeVendorApiListRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeVendorApiListRequest) SetPageNumber(v int32) *DescribeVendorApiListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVendorApiListRequest) SetPageSize(v int64) *DescribeVendorApiListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVendorApiListRequest) SetProductCode(v string) *DescribeVendorApiListRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeVendorApiListRequest) SetVendorCode(v string) *DescribeVendorApiListRequest {
	s.VendorCode = &v
	return s
}

type DescribeVendorApiListResponseBody struct {
	ApiList []*DescribeVendorApiListResponseBodyApiList `json:"ApiList,omitempty" xml:"ApiList,omitempty" type:"Repeated"`
	Page    *DescribeVendorApiListResponseBodyPage      `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// E7698CFB-****-5840-8EC9-691B86729E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVendorApiListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVendorApiListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVendorApiListResponseBody) SetApiList(v []*DescribeVendorApiListResponseBodyApiList) *DescribeVendorApiListResponseBody {
	s.ApiList = v
	return s
}

func (s *DescribeVendorApiListResponseBody) SetPage(v *DescribeVendorApiListResponseBodyPage) *DescribeVendorApiListResponseBody {
	s.Page = v
	return s
}

func (s *DescribeVendorApiListResponseBody) SetRequestId(v string) *DescribeVendorApiListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVendorApiListResponseBodyApiList struct {
	// example:
	//
	// {
	//
	//     "cmd": "DescribeAclApiDispatch"
	//
	// }
	AdvanceConfig *string `json:"AdvanceConfig,omitempty" xml:"AdvanceConfig,omitempty"`
	// example:
	//
	// VerifyMobile
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// example:
	//
	// 2017-08-01
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// example:
	//
	// true
	NeedAdvanceConfig *bool `json:"NeedAdvanceConfig,omitempty" xml:"NeedAdvanceConfig,omitempty"`
	// example:
	//
	// false
	NeedPageInfo *bool `json:"NeedPageInfo,omitempty" xml:"NeedPageInfo,omitempty"`
	// example:
	//
	// {\\"Count\\": 10, \\"TotalCount\\": 23, \\"PageSize\\": 10, \\"CurrentPage\\": 1}
	PageInfo *string `json:"PageInfo,omitempty" xml:"PageInfo,omitempty"`
	// example:
	//
	// [
	//
	//     {
	//
	//         "name": "Domain",
	//
	//         "type": "String",
	//
	//         "isRequired": true,
	//
	//         "exampleValue": "www.***.com",
	//
	//         "description": "www.***.com"
	//
	//     }
	//
	// ]
	Parameter *string `json:"Parameter,omitempty" xml:"Parameter,omitempty"`
	// example:
	//
	// cfw
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// cfw.xxx.com
	ProductDomain *string `json:"ProductDomain,omitempty" xml:"ProductDomain,omitempty"`
	// example:
	//
	// waf
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// https
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// Azure
	VendorCode *string `json:"VendorCode,omitempty" xml:"VendorCode,omitempty"`
}

func (s DescribeVendorApiListResponseBodyApiList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVendorApiListResponseBodyApiList) GoString() string {
	return s.String()
}

func (s *DescribeVendorApiListResponseBodyApiList) SetAdvanceConfig(v string) *DescribeVendorApiListResponseBodyApiList {
	s.AdvanceConfig = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyApiList) SetApiName(v string) *DescribeVendorApiListResponseBodyApiList {
	s.ApiName = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyApiList) SetApiVersion(v string) *DescribeVendorApiListResponseBodyApiList {
	s.ApiVersion = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyApiList) SetMethod(v string) *DescribeVendorApiListResponseBodyApiList {
	s.Method = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyApiList) SetNeedAdvanceConfig(v bool) *DescribeVendorApiListResponseBodyApiList {
	s.NeedAdvanceConfig = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyApiList) SetNeedPageInfo(v bool) *DescribeVendorApiListResponseBodyApiList {
	s.NeedPageInfo = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyApiList) SetPageInfo(v string) *DescribeVendorApiListResponseBodyApiList {
	s.PageInfo = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyApiList) SetParameter(v string) *DescribeVendorApiListResponseBodyApiList {
	s.Parameter = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyApiList) SetProductCode(v string) *DescribeVendorApiListResponseBodyApiList {
	s.ProductCode = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyApiList) SetProductDomain(v string) *DescribeVendorApiListResponseBodyApiList {
	s.ProductDomain = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyApiList) SetProductName(v string) *DescribeVendorApiListResponseBodyApiList {
	s.ProductName = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyApiList) SetProtocol(v string) *DescribeVendorApiListResponseBodyApiList {
	s.Protocol = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyApiList) SetVendorCode(v string) *DescribeVendorApiListResponseBodyApiList {
	s.VendorCode = &v
	return s
}

type DescribeVendorApiListResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVendorApiListResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s DescribeVendorApiListResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeVendorApiListResponseBodyPage) SetPageNumber(v int64) *DescribeVendorApiListResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyPage) SetPageSize(v int32) *DescribeVendorApiListResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeVendorApiListResponseBodyPage) SetTotalCount(v int64) *DescribeVendorApiListResponseBodyPage {
	s.TotalCount = &v
	return s
}

type DescribeVendorApiListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVendorApiListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVendorApiListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVendorApiListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVendorApiListResponse) SetHeaders(v map[string]*string) *DescribeVendorApiListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVendorApiListResponse) SetStatusCode(v int32) *DescribeVendorApiListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVendorApiListResponse) SetBody(v *DescribeVendorApiListResponseBody) *DescribeVendorApiListResponse {
	s.Body = v
	return s
}

type DescriberPython3ScriptLogsRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID that is returned when the Python3 script is run.
	//
	// >  You can call the [RunPython3Script](~~RunPython3Script~~) operation to query the UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 69edc2b4-c95c-424f-9114-xxxxxxx
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
}

func (s DescriberPython3ScriptLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescriberPython3ScriptLogsRequest) GoString() string {
	return s.String()
}

func (s *DescriberPython3ScriptLogsRequest) SetLang(v string) *DescriberPython3ScriptLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescriberPython3ScriptLogsRequest) SetRequestUuid(v string) *DescriberPython3ScriptLogsRequest {
	s.RequestUuid = &v
	return s
}

type DescriberPython3ScriptLogsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D22D8A0C-6E86-57B2-A142-929184122AB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The operational logs of the Python3 script.
	//
	// example:
	//
	// {
	//
	//     "logs": [
	//
	//         {
	//
	//             "message": "function input is {}"
	//
	//         }
	//
	//     ]
	//
	// }
	RunResult *string `json:"RunResult,omitempty" xml:"RunResult,omitempty"`
}

func (s DescriberPython3ScriptLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescriberPython3ScriptLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescriberPython3ScriptLogsResponseBody) SetRequestId(v string) *DescriberPython3ScriptLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescriberPython3ScriptLogsResponseBody) SetRunResult(v string) *DescriberPython3ScriptLogsResponseBody {
	s.RunResult = &v
	return s
}

type DescriberPython3ScriptLogsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescriberPython3ScriptLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescriberPython3ScriptLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescriberPython3ScriptLogsResponse) GoString() string {
	return s.String()
}

func (s *DescriberPython3ScriptLogsResponse) SetHeaders(v map[string]*string) *DescriberPython3ScriptLogsResponse {
	s.Headers = v
	return s
}

func (s *DescriberPython3ScriptLogsResponse) SetStatusCode(v int32) *DescriberPython3ScriptLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescriberPython3ScriptLogsResponse) SetBody(v *DescriberPython3ScriptLogsResponseBody) *DescriberPython3ScriptLogsResponse {
	s.Body = v
	return s
}

type ModifyComponentAssetRequest struct {
	// The configuration of the asset. The value is a JSON object.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "name": "test asset",
	//
	//     "componentName": "SLS",
	//
	//     "params": [
	//
	//         {
	//
	//             "name": "end_point",
	//
	//             "value": "xxx"
	//
	//         },
	//
	//         {
	//
	//             "name": "sub_id",
	//
	//             "value": "xxxx"
	//
	//         },
	//
	//         {
	//
	//             "name": "access_key",
	//
	//             "value": "xxxx"
	//
	//         }
	//
	//     ]
	//
	// }
	AssetConfig *string `json:"AssetConfig,omitempty" xml:"AssetConfig,omitempty"`
	// The language of the content within the request and response.
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ModifyComponentAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyComponentAssetRequest) GoString() string {
	return s.String()
}

func (s *ModifyComponentAssetRequest) SetAssetConfig(v string) *ModifyComponentAssetRequest {
	s.AssetConfig = &v
	return s
}

func (s *ModifyComponentAssetRequest) SetLang(v string) *ModifyComponentAssetRequest {
	s.Lang = &v
	return s
}

type ModifyComponentAssetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1C5F11E9-464E-51F0-9296-43BB312A0557
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyComponentAssetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyComponentAssetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyComponentAssetResponseBody) SetRequestId(v string) *ModifyComponentAssetResponseBody {
	s.RequestId = &v
	return s
}

type ModifyComponentAssetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyComponentAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyComponentAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyComponentAssetResponse) GoString() string {
	return s.String()
}

func (s *ModifyComponentAssetResponse) SetHeaders(v map[string]*string) *ModifyComponentAssetResponse {
	s.Headers = v
	return s
}

func (s *ModifyComponentAssetResponse) SetStatusCode(v int32) *ModifyComponentAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyComponentAssetResponse) SetBody(v *ModifyComponentAssetResponseBody) *ModifyComponentAssetResponse {
	s.Body = v
	return s
}

type ModifyPlaybookRequest struct {
	// The description of the playbook.
	//
	// example:
	//
	// demo test task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun_waf_test_playbook
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-1586c35e61f8
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The XML configuration of the playbook.
	//
	// example:
	//
	// <?xml version="1.0" encoding="UTF-8"?><bpmn:definitions xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:bpmn="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:dc="http://www.omg.org/spec/DD/20100524/DC" targetNamespace="http://bpmn.io/schema/bpmn" id="Definitions_1"><bpmn:process id="Process_1" isExecutable="false"><bpmn:startEvent id="StartEvent_1"/></bpmn:process><bpmndi:BPMNDiagram id="BPMNDiagram_1"><bpmndi:BPMNPlane id="BPMNPlane_1" bpmnElement="Process_1"><bpmndi:BPMNShape id="_BPMNShape_StartEvent_2" bpmnElement="StartEvent_1"><dc:Bounds height="36.0" width="36.0" x="173.0" y="102.0"/></bpmndi:BPMNShape></bpmndi:BPMNPlane></bpmndi:BPMNDiagram></bpmn:definitions>
	Taskflow *string `json:"Taskflow,omitempty" xml:"Taskflow,omitempty"`
}

func (s ModifyPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookRequest) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookRequest) SetDescription(v string) *ModifyPlaybookRequest {
	s.Description = &v
	return s
}

func (s *ModifyPlaybookRequest) SetDisplayName(v string) *ModifyPlaybookRequest {
	s.DisplayName = &v
	return s
}

func (s *ModifyPlaybookRequest) SetLang(v string) *ModifyPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPlaybookRequest) SetPlaybookUuid(v string) *ModifyPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ModifyPlaybookRequest) SetTaskflow(v string) *ModifyPlaybookRequest {
	s.Taskflow = &v
	return s
}

type ModifyPlaybookResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B584F84-D66A-5525-8E7B-05612A903ABF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookResponseBody) SetRequestId(v string) *ModifyPlaybookResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookResponse) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookResponse) SetHeaders(v map[string]*string) *ModifyPlaybookResponse {
	s.Headers = v
	return s
}

func (s *ModifyPlaybookResponse) SetStatusCode(v int32) *ModifyPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPlaybookResponse) SetBody(v *ModifyPlaybookResponseBody) *ModifyPlaybookResponse {
	s.Body = v
	return s
}

type ModifyPlaybookInputOutputRequest struct {
	// The executed mode of a playbook. The value is a JSON array.
	ExeConfig *string `json:"ExeConfig,omitempty" xml:"ExeConfig,omitempty"`
	// The configuration of the input parameters. The value is a JSON array.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "typeName": "String",
	//
	//         "dataClass": "normal",
	//
	//         "dataType": "String",
	//
	//         "description": "period",
	//
	//         "example": "",
	//
	//         "name": "period",
	//
	//         "required": false
	//
	//     }
	//
	// ]
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// The language of the content within the request and response.
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The configuration of the output parameters. This parameter is unavailable. Leave it empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// []
	OutputParams *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	// The input parameter type.
	//
	// 	- **template-ip**
	//
	// 	- **template-file**
	//
	// 	- **template-process**
	//
	// 	- **custom**
	//
	// example:
	//
	// custom
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s ModifyPlaybookInputOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookInputOutputRequest) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInputOutputRequest) SetExeConfig(v string) *ModifyPlaybookInputOutputRequest {
	s.ExeConfig = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetInputParams(v string) *ModifyPlaybookInputOutputRequest {
	s.InputParams = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetLang(v string) *ModifyPlaybookInputOutputRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetOutputParams(v string) *ModifyPlaybookInputOutputRequest {
	s.OutputParams = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetParamType(v string) *ModifyPlaybookInputOutputRequest {
	s.ParamType = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetPlaybookUuid(v string) *ModifyPlaybookInputOutputRequest {
	s.PlaybookUuid = &v
	return s
}

type ModifyPlaybookInputOutputResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8DDC07CE-D41B-5142-8D91-469462719C77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPlaybookInputOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookInputOutputResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInputOutputResponseBody) SetRequestId(v string) *ModifyPlaybookInputOutputResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPlaybookInputOutputResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPlaybookInputOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPlaybookInputOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookInputOutputResponse) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInputOutputResponse) SetHeaders(v map[string]*string) *ModifyPlaybookInputOutputResponse {
	s.Headers = v
	return s
}

func (s *ModifyPlaybookInputOutputResponse) SetStatusCode(v int32) *ModifyPlaybookInputOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPlaybookInputOutputResponse) SetBody(v *ModifyPlaybookInputOutputResponseBody) *ModifyPlaybookInputOutputResponse {
	s.Body = v
	return s
}

type ModifyPlaybookInstanceStatusRequest struct {
	// The playbook status. Valid values:
	//
	// 	- **1**: starts the playbook.
	//
	// 	- **0**: stops the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Active *int32 `json:"Active,omitempty" xml:"Active,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9fcd3829-80ff-4681-be1e-xxxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s ModifyPlaybookInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInstanceStatusRequest) SetActive(v int32) *ModifyPlaybookInstanceStatusRequest {
	s.Active = &v
	return s
}

func (s *ModifyPlaybookInstanceStatusRequest) SetLang(v string) *ModifyPlaybookInstanceStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPlaybookInstanceStatusRequest) SetPlaybookUuid(v string) *ModifyPlaybookInstanceStatusRequest {
	s.PlaybookUuid = &v
	return s
}

type ModifyPlaybookInstanceStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C2A32830-2842-5F8F-B4ED-E4783E400BBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPlaybookInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInstanceStatusResponseBody) SetRequestId(v string) *ModifyPlaybookInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPlaybookInstanceStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPlaybookInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPlaybookInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPlaybookInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInstanceStatusResponse) SetHeaders(v map[string]*string) *ModifyPlaybookInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyPlaybookInstanceStatusResponse) SetStatusCode(v int32) *ModifyPlaybookInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPlaybookInstanceStatusResponse) SetBody(v *ModifyPlaybookInstanceStatusResponseBody) *ModifyPlaybookInstanceStatusResponse {
	s.Body = v
	return s
}

type PublishPlaybookRequest struct {
	// The description of the released version.
	//
	// example:
	//
	// This is a waf processing playbook
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac343acc-1a61-4084-9a1c-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s PublishPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishPlaybookRequest) GoString() string {
	return s.String()
}

func (s *PublishPlaybookRequest) SetDescription(v string) *PublishPlaybookRequest {
	s.Description = &v
	return s
}

func (s *PublishPlaybookRequest) SetPlaybookUuid(v string) *PublishPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

type PublishPlaybookResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C513FCEA-D71F-5E50-ADC4-FCF8C5DCF6BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *PublishPlaybookResponseBody) SetRequestId(v string) *PublishPlaybookResponseBody {
	s.RequestId = &v
	return s
}

type PublishPlaybookResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishPlaybookResponse) GoString() string {
	return s.String()
}

func (s *PublishPlaybookResponse) SetHeaders(v map[string]*string) *PublishPlaybookResponse {
	s.Headers = v
	return s
}

func (s *PublishPlaybookResponse) SetStatusCode(v int32) *PublishPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishPlaybookResponse) SetBody(v *PublishPlaybookResponseBody) *PublishPlaybookResponse {
	s.Body = v
	return s
}

type QueryTreeDataRequest struct {
	// The language of the content within the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s QueryTreeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTreeDataRequest) GoString() string {
	return s.String()
}

func (s *QueryTreeDataRequest) SetLang(v string) *QueryTreeDataRequest {
	s.Lang = &v
	return s
}

type QueryTreeDataResponseBody struct {
	// The returned information about the playbook. The value is a JSON string.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "playbook": {
	//
	//             "active": false,
	//
	//             "displayName": "test_playbook",
	//
	//             "playbookUuid": "09a20455-3d3a-424c-a1df-xxxxxx"
	//
	//         }
	//
	//     }
	//
	// ]
	Playbooks *string `json:"Playbooks,omitempty" xml:"Playbooks,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EF2ECA2D-D8E6-5021-BF5C-19DD6D52C5B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTreeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTreeDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTreeDataResponseBody) SetPlaybooks(v string) *QueryTreeDataResponseBody {
	s.Playbooks = &v
	return s
}

func (s *QueryTreeDataResponseBody) SetRequestId(v string) *QueryTreeDataResponseBody {
	s.RequestId = &v
	return s
}

type QueryTreeDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTreeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTreeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTreeDataResponse) GoString() string {
	return s.String()
}

func (s *QueryTreeDataResponse) SetHeaders(v map[string]*string) *QueryTreeDataResponse {
	s.Headers = v
	return s
}

func (s *QueryTreeDataResponse) SetStatusCode(v int32) *QueryTreeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTreeDataResponse) SetBody(v *QueryTreeDataResponseBody) *QueryTreeDataResponse {
	s.Body = v
	return s
}

type RenamePlaybookNodeRequest struct {
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The new name of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_process
	NewNodeName *string `json:"NewNodeName,omitempty" xml:"NewNodeName,omitempty"`
	// The original name of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// firewall_process
	OldNodeName *string `json:"OldNodeName,omitempty" xml:"OldNodeName,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac343acc-1a61-4084-9a1c-xxxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s RenamePlaybookNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RenamePlaybookNodeRequest) GoString() string {
	return s.String()
}

func (s *RenamePlaybookNodeRequest) SetLang(v string) *RenamePlaybookNodeRequest {
	s.Lang = &v
	return s
}

func (s *RenamePlaybookNodeRequest) SetNewNodeName(v string) *RenamePlaybookNodeRequest {
	s.NewNodeName = &v
	return s
}

func (s *RenamePlaybookNodeRequest) SetOldNodeName(v string) *RenamePlaybookNodeRequest {
	s.OldNodeName = &v
	return s
}

func (s *RenamePlaybookNodeRequest) SetPlaybookUuid(v string) *RenamePlaybookNodeRequest {
	s.PlaybookUuid = &v
	return s
}

type RenamePlaybookNodeResponseBody struct {
	// The returned new name of the node.
	//
	// example:
	//
	// waf_process
	RenameResult *string `json:"RenameResult,omitempty" xml:"RenameResult,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1E1EC464-3BD7-518F-9937-BCC12E6855FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenamePlaybookNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenamePlaybookNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RenamePlaybookNodeResponseBody) SetRenameResult(v string) *RenamePlaybookNodeResponseBody {
	s.RenameResult = &v
	return s
}

func (s *RenamePlaybookNodeResponseBody) SetRequestId(v string) *RenamePlaybookNodeResponseBody {
	s.RequestId = &v
	return s
}

type RenamePlaybookNodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenamePlaybookNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenamePlaybookNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RenamePlaybookNodeResponse) GoString() string {
	return s.String()
}

func (s *RenamePlaybookNodeResponse) SetHeaders(v map[string]*string) *RenamePlaybookNodeResponse {
	s.Headers = v
	return s
}

func (s *RenamePlaybookNodeResponse) SetStatusCode(v int32) *RenamePlaybookNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RenamePlaybookNodeResponse) SetBody(v *RenamePlaybookNodeResponseBody) *RenamePlaybookNodeResponse {
	s.Body = v
	return s
}

type RevertPlaybookReleaseRequest struct {
	// Specifies whether to directly publish the new playbook after the rollback.
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsPublish *bool `json:"IsPublish,omitempty" xml:"IsPublish,omitempty"`
	// The version of the playbook that you want to publish.
	//
	// >  You can call the [DescribePlaybookReleases](~~DescribePlaybookReleases~~) operation to query the playbook version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3f97b56e-064e-47e7-a309-xxxxxxx
	PlayReleaseId *int32 `json:"PlayReleaseId,omitempty" xml:"PlayReleaseId,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 185295a1-c987-4b64-8796-xxxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s RevertPlaybookReleaseRequest) String() string {
	return tea.Prettify(s)
}

func (s RevertPlaybookReleaseRequest) GoString() string {
	return s.String()
}

func (s *RevertPlaybookReleaseRequest) SetIsPublish(v bool) *RevertPlaybookReleaseRequest {
	s.IsPublish = &v
	return s
}

func (s *RevertPlaybookReleaseRequest) SetPlayReleaseId(v int32) *RevertPlaybookReleaseRequest {
	s.PlayReleaseId = &v
	return s
}

func (s *RevertPlaybookReleaseRequest) SetPlaybookUuid(v string) *RevertPlaybookReleaseRequest {
	s.PlaybookUuid = &v
	return s
}

type RevertPlaybookReleaseResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B3FED5B9-190A-5952-93A4-24FBF0F0C573
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevertPlaybookReleaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevertPlaybookReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *RevertPlaybookReleaseResponseBody) SetRequestId(v string) *RevertPlaybookReleaseResponseBody {
	s.RequestId = &v
	return s
}

type RevertPlaybookReleaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevertPlaybookReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertPlaybookReleaseResponse) String() string {
	return tea.Prettify(s)
}

func (s RevertPlaybookReleaseResponse) GoString() string {
	return s.String()
}

func (s *RevertPlaybookReleaseResponse) SetHeaders(v map[string]*string) *RevertPlaybookReleaseResponse {
	s.Headers = v
	return s
}

func (s *RevertPlaybookReleaseResponse) SetStatusCode(v int32) *RevertPlaybookReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertPlaybookReleaseResponse) SetBody(v *RevertPlaybookReleaseResponseBody) *RevertPlaybookReleaseResponse {
	s.Body = v
	return s
}

type RunNotifyComponentWithEmailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// notifyByCustom
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	AssetId *int32 `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NotifyMessage
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// email content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// notify_message_1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e99dab31-499b-4307-9248-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// This parameter is required.
	Receivers []*string `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
	// example:
	//
	// 137602xxx718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// title
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s RunNotifyComponentWithEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s RunNotifyComponentWithEmailRequest) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithEmailRequest) SetActionName(v string) *RunNotifyComponentWithEmailRequest {
	s.ActionName = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetAssetId(v int32) *RunNotifyComponentWithEmailRequest {
	s.AssetId = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetComponentName(v string) *RunNotifyComponentWithEmailRequest {
	s.ComponentName = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetContent(v string) *RunNotifyComponentWithEmailRequest {
	s.Content = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetLang(v string) *RunNotifyComponentWithEmailRequest {
	s.Lang = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetNodeName(v string) *RunNotifyComponentWithEmailRequest {
	s.NodeName = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetPlaybookUuid(v string) *RunNotifyComponentWithEmailRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetReceivers(v []*string) *RunNotifyComponentWithEmailRequest {
	s.Receivers = v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetRoleFor(v int64) *RunNotifyComponentWithEmailRequest {
	s.RoleFor = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetRoleType(v string) *RunNotifyComponentWithEmailRequest {
	s.RoleType = &v
	return s
}

func (s *RunNotifyComponentWithEmailRequest) SetSubject(v string) *RunNotifyComponentWithEmailRequest {
	s.Subject = &v
	return s
}

type RunNotifyComponentWithEmailResponseBody struct {
	// example:
	//
	// {}
	Data *string                                      `json:"Data,omitempty" xml:"Data,omitempty"`
	Page *RunNotifyComponentWithEmailResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// D4CC979E-3D5B-5A6A-BC87-C93C9E861C7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunNotifyComponentWithEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunNotifyComponentWithEmailResponseBody) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithEmailResponseBody) SetData(v string) *RunNotifyComponentWithEmailResponseBody {
	s.Data = &v
	return s
}

func (s *RunNotifyComponentWithEmailResponseBody) SetPage(v *RunNotifyComponentWithEmailResponseBodyPage) *RunNotifyComponentWithEmailResponseBody {
	s.Page = v
	return s
}

func (s *RunNotifyComponentWithEmailResponseBody) SetRequestId(v string) *RunNotifyComponentWithEmailResponseBody {
	s.RequestId = &v
	return s
}

type RunNotifyComponentWithEmailResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s RunNotifyComponentWithEmailResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s RunNotifyComponentWithEmailResponseBodyPage) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithEmailResponseBodyPage) SetPageNumber(v int32) *RunNotifyComponentWithEmailResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *RunNotifyComponentWithEmailResponseBodyPage) SetPageSize(v int32) *RunNotifyComponentWithEmailResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *RunNotifyComponentWithEmailResponseBodyPage) SetTotalCount(v int32) *RunNotifyComponentWithEmailResponseBodyPage {
	s.TotalCount = &v
	return s
}

type RunNotifyComponentWithEmailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunNotifyComponentWithEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunNotifyComponentWithEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s RunNotifyComponentWithEmailResponse) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithEmailResponse) SetHeaders(v map[string]*string) *RunNotifyComponentWithEmailResponse {
	s.Headers = v
	return s
}

func (s *RunNotifyComponentWithEmailResponse) SetStatusCode(v int32) *RunNotifyComponentWithEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *RunNotifyComponentWithEmailResponse) SetBody(v *RunNotifyComponentWithEmailResponseBody) *RunNotifyComponentWithEmailResponse {
	s.Body = v
	return s
}

type RunNotifyComponentWithMessageCenterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// notifyByMessageCenter
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 146789xxxx733152
	Aliuid *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// example:
	//
	// 1
	AssetId         *int32    `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	ChannelTypeList []*string `json:"ChannelTypeList,omitempty" xml:"ChannelTypeList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// NotifyMessage
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yundun_soar_incident_generate
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// notify_message
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// {"startTime":"test222","incidentName":"test123","incidentID":"teset123"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c5c88b5e-97ca-435d-8c20-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// example:
	//
	// 1467894xxx733152
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s RunNotifyComponentWithMessageCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s RunNotifyComponentWithMessageCenterRequest) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetActionName(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.ActionName = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetAliuid(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.Aliuid = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetAssetId(v int32) *RunNotifyComponentWithMessageCenterRequest {
	s.AssetId = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetChannelTypeList(v []*string) *RunNotifyComponentWithMessageCenterRequest {
	s.ChannelTypeList = v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetComponentName(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.ComponentName = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetEventId(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.EventId = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetLang(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.Lang = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetNodeName(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.NodeName = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetParams(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.Params = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetPlaybookUuid(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetRoleFor(v int64) *RunNotifyComponentWithMessageCenterRequest {
	s.RoleFor = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterRequest) SetRoleType(v string) *RunNotifyComponentWithMessageCenterRequest {
	s.RoleType = &v
	return s
}

type RunNotifyComponentWithMessageCenterResponseBody struct {
	// example:
	//
	// {}
	Data *string                                              `json:"Data,omitempty" xml:"Data,omitempty"`
	Page *RunNotifyComponentWithMessageCenterResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// E7698CFB-4E1C-5840-8EC9-691B86729E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunNotifyComponentWithMessageCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunNotifyComponentWithMessageCenterResponseBody) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithMessageCenterResponseBody) SetData(v string) *RunNotifyComponentWithMessageCenterResponseBody {
	s.Data = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponseBody) SetPage(v *RunNotifyComponentWithMessageCenterResponseBodyPage) *RunNotifyComponentWithMessageCenterResponseBody {
	s.Page = v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponseBody) SetRequestId(v string) *RunNotifyComponentWithMessageCenterResponseBody {
	s.RequestId = &v
	return s
}

type RunNotifyComponentWithMessageCenterResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s RunNotifyComponentWithMessageCenterResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s RunNotifyComponentWithMessageCenterResponseBodyPage) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithMessageCenterResponseBodyPage) SetPageNumber(v int32) *RunNotifyComponentWithMessageCenterResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponseBodyPage) SetPageSize(v int32) *RunNotifyComponentWithMessageCenterResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponseBodyPage) SetTotalCount(v int32) *RunNotifyComponentWithMessageCenterResponseBodyPage {
	s.TotalCount = &v
	return s
}

type RunNotifyComponentWithMessageCenterResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunNotifyComponentWithMessageCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunNotifyComponentWithMessageCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s RunNotifyComponentWithMessageCenterResponse) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithMessageCenterResponse) SetHeaders(v map[string]*string) *RunNotifyComponentWithMessageCenterResponse {
	s.Headers = v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponse) SetStatusCode(v int32) *RunNotifyComponentWithMessageCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponse) SetBody(v *RunNotifyComponentWithMessageCenterResponseBody) *RunNotifyComponentWithMessageCenterResponse {
	s.Body = v
	return s
}

type RunNotifyComponentWithWebhookRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// notifyByCustom
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// example:
	//
	// 1
	AssetId *int32 `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NotifyMessage
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "at": {
	//
	//         "atMobiles":[
	//
	//             "180xxxxxx"
	//
	//         ],
	//
	//         "atUserIds":[
	//
	//             "user123"
	//
	//         ],
	//
	//         "isAtAll": false
	//
	//     },
	//
	//     "text": {
	//
	//         "content":"1234"
	//
	//     },
	//
	//     "msgtype":"text"
	//
	// }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	MsgType *string `json:"MsgType,omitempty" xml:"MsgType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// notify_message_node
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 94bc318c-****-4cba-****-801ccb0d739f
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// example:
	//
	// 126339xxxx805497
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// example:
	//
	// SECc1*****e157b32b380f********bb8c70e1a67a22072
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [\\"10651\\"]
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s RunNotifyComponentWithWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s RunNotifyComponentWithWebhookRequest) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithWebhookRequest) SetActionName(v string) *RunNotifyComponentWithWebhookRequest {
	s.ActionName = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetAssetId(v int32) *RunNotifyComponentWithWebhookRequest {
	s.AssetId = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetComponentName(v string) *RunNotifyComponentWithWebhookRequest {
	s.ComponentName = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetContent(v string) *RunNotifyComponentWithWebhookRequest {
	s.Content = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetLang(v string) *RunNotifyComponentWithWebhookRequest {
	s.Lang = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetMsgType(v string) *RunNotifyComponentWithWebhookRequest {
	s.MsgType = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetNodeName(v string) *RunNotifyComponentWithWebhookRequest {
	s.NodeName = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetPlaybookUuid(v string) *RunNotifyComponentWithWebhookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetRoleFor(v int64) *RunNotifyComponentWithWebhookRequest {
	s.RoleFor = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetRoleType(v string) *RunNotifyComponentWithWebhookRequest {
	s.RoleType = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetSecret(v string) *RunNotifyComponentWithWebhookRequest {
	s.Secret = &v
	return s
}

func (s *RunNotifyComponentWithWebhookRequest) SetWebhook(v string) *RunNotifyComponentWithWebhookRequest {
	s.Webhook = &v
	return s
}

type RunNotifyComponentWithWebhookResponseBody struct {
	// example:
	//
	// {}
	Data *string                                        `json:"Data,omitempty" xml:"Data,omitempty"`
	Page *RunNotifyComponentWithWebhookResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// E7698CFB-****-5840-8EC9-691B86729E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunNotifyComponentWithWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunNotifyComponentWithWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithWebhookResponseBody) SetData(v string) *RunNotifyComponentWithWebhookResponseBody {
	s.Data = &v
	return s
}

func (s *RunNotifyComponentWithWebhookResponseBody) SetPage(v *RunNotifyComponentWithWebhookResponseBodyPage) *RunNotifyComponentWithWebhookResponseBody {
	s.Page = v
	return s
}

func (s *RunNotifyComponentWithWebhookResponseBody) SetRequestId(v string) *RunNotifyComponentWithWebhookResponseBody {
	s.RequestId = &v
	return s
}

type RunNotifyComponentWithWebhookResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s RunNotifyComponentWithWebhookResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s RunNotifyComponentWithWebhookResponseBodyPage) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithWebhookResponseBodyPage) SetPageNumber(v int32) *RunNotifyComponentWithWebhookResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *RunNotifyComponentWithWebhookResponseBodyPage) SetPageSize(v int32) *RunNotifyComponentWithWebhookResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *RunNotifyComponentWithWebhookResponseBodyPage) SetTotalCount(v int32) *RunNotifyComponentWithWebhookResponseBodyPage {
	s.TotalCount = &v
	return s
}

type RunNotifyComponentWithWebhookResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunNotifyComponentWithWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunNotifyComponentWithWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s RunNotifyComponentWithWebhookResponse) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithWebhookResponse) SetHeaders(v map[string]*string) *RunNotifyComponentWithWebhookResponse {
	s.Headers = v
	return s
}

func (s *RunNotifyComponentWithWebhookResponse) SetStatusCode(v int32) *RunNotifyComponentWithWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *RunNotifyComponentWithWebhookResponse) SetBody(v *RunNotifyComponentWithWebhookResponseBody) *RunNotifyComponentWithWebhookResponse {
	s.Body = v
	return s
}

type RunPython3ScriptRequest struct {
	// The name of the node in the playbook.
	//
	// example:
	//
	// python3_3
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The input parameters of the Python3 script.
	//
	// example:
	//
	// {
	//
	//     "input1": "xx.xx.xx.xx",
	//
	//     "input2": "7d"
	//
	// }
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the UUIDs of playbooks.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The Python3 script.
	//
	// example:
	//
	// import logging
	//
	// def execute (params):
	//
	//   #ip = params[\\"ip\\"]
	//
	//   #logging.info("enter execute,ip is "+ip)
	//
	//   success=True
	//
	//   message=\\"OK\\"
	//
	//   data=[]
	//
	//   return (success,message,data)
	PythonScript *string `json:"PythonScript,omitempty" xml:"PythonScript,omitempty"`
}

func (s RunPython3ScriptRequest) String() string {
	return tea.Prettify(s)
}

func (s RunPython3ScriptRequest) GoString() string {
	return s.String()
}

func (s *RunPython3ScriptRequest) SetNodeName(v string) *RunPython3ScriptRequest {
	s.NodeName = &v
	return s
}

func (s *RunPython3ScriptRequest) SetParams(v string) *RunPython3ScriptRequest {
	s.Params = &v
	return s
}

func (s *RunPython3ScriptRequest) SetPlaybookUuid(v string) *RunPython3ScriptRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *RunPython3ScriptRequest) SetPythonScript(v string) *RunPython3ScriptRequest {
	s.PythonScript = &v
	return s
}

type RunPython3ScriptResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F210521C-D9BF-5264-8369-83EDDC617DB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the Python3 script.
	//
	// example:
	//
	// {
	//
	//     "requestUuid": "fe240b98-27b1-4a36-aec1-550b894318d9",
	//
	//     "content": {
	//
	//         "resultData": [],
	//
	//         "success": true
	//
	//     }
	//
	// }
	RunResult *string `json:"RunResult,omitempty" xml:"RunResult,omitempty"`
}

func (s RunPython3ScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunPython3ScriptResponseBody) GoString() string {
	return s.String()
}

func (s *RunPython3ScriptResponseBody) SetRequestId(v string) *RunPython3ScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunPython3ScriptResponseBody) SetRunResult(v string) *RunPython3ScriptResponseBody {
	s.RunResult = &v
	return s
}

type RunPython3ScriptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunPython3ScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunPython3ScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s RunPython3ScriptResponse) GoString() string {
	return s.String()
}

func (s *RunPython3ScriptResponse) SetHeaders(v map[string]*string) *RunPython3ScriptResponse {
	s.Headers = v
	return s
}

func (s *RunPython3ScriptResponse) SetStatusCode(v int32) *RunPython3ScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *RunPython3ScriptResponse) SetBody(v *RunPython3ScriptResponseBody) *RunPython3ScriptResponse {
	s.Body = v
	return s
}

type TriggerPlaybookRequest struct {
	// The input parameters of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "input1": "xx.xx.xx.xx",
	//
	//     "input2": "7d"
	//
	// }
	InputParam *string `json:"InputParam,omitempty" xml:"InputParam,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2a687089-d4dd-47d4-9709-xxxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s TriggerPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerPlaybookRequest) GoString() string {
	return s.String()
}

func (s *TriggerPlaybookRequest) SetInputParam(v string) *TriggerPlaybookRequest {
	s.InputParam = &v
	return s
}

func (s *TriggerPlaybookRequest) SetPlaybookUuid(v string) *TriggerPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

type TriggerPlaybookResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD5A8DB6-A42C-532B-BCE8-83E69550CD59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The running UUID of the playbook. This parameter is used to query the running result of the playbook.
	//
	// example:
	//
	// 55E63C57-D6C8-5036-A770-5CB10AC807AA
	TriggerUuid *string `json:"TriggerUuid,omitempty" xml:"TriggerUuid,omitempty"`
}

func (s TriggerPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerPlaybookResponseBody) SetRequestId(v string) *TriggerPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerPlaybookResponseBody) SetTriggerUuid(v string) *TriggerPlaybookResponseBody {
	s.TriggerUuid = &v
	return s
}

type TriggerPlaybookResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerPlaybookResponse) GoString() string {
	return s.String()
}

func (s *TriggerPlaybookResponse) SetHeaders(v map[string]*string) *TriggerPlaybookResponse {
	s.Headers = v
	return s
}

func (s *TriggerPlaybookResponse) SetStatusCode(v int32) *TriggerPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerPlaybookResponse) SetBody(v *TriggerPlaybookResponseBody) *TriggerPlaybookResponse {
	s.Body = v
	return s
}

type TriggerProcessTaskRequest struct {
	// The type of the action. Valid values:
	//
	// 	- **remove**: cancels blocking or isolation.
	//
	// 	- **retry**: submits the task again.
	//
	// This parameter is required.
	//
	// example:
	//
	// remove
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The ID of the handling task.
	//
	// >  You can call the [DescribeProcessTasks](~~DescribeProcessTasks~~) operation to query the IDs of handling tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15355xxxxxx82894882
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TriggerProcessTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerProcessTaskRequest) GoString() string {
	return s.String()
}

func (s *TriggerProcessTaskRequest) SetActionType(v string) *TriggerProcessTaskRequest {
	s.ActionType = &v
	return s
}

func (s *TriggerProcessTaskRequest) SetTaskId(v string) *TriggerProcessTaskRequest {
	s.TaskId = &v
	return s
}

type TriggerProcessTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 58A518BC-E4A8-5BD7-AFEA-366046ED9073
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerProcessTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerProcessTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerProcessTaskResponseBody) SetRequestId(v string) *TriggerProcessTaskResponseBody {
	s.RequestId = &v
	return s
}

type TriggerProcessTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerProcessTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerProcessTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerProcessTaskResponse) GoString() string {
	return s.String()
}

func (s *TriggerProcessTaskResponse) SetHeaders(v map[string]*string) *TriggerProcessTaskResponse {
	s.Headers = v
	return s
}

func (s *TriggerProcessTaskResponse) SetStatusCode(v int32) *TriggerProcessTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerProcessTaskResponse) SetBody(v *TriggerProcessTaskResponseBody) *TriggerProcessTaskResponse {
	s.Body = v
	return s
}

type TriggerSophonPlaybookRequest struct {
	// The name of the command that you want to trigger.
	//
	// >  You can call the [DescribeSophonCommands](~~DescribeSophonCommands~~) operation to query the command name.
	//
	// example:
	//
	// waf_process_command
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The input parameters of the command or playbook that you want to trigger.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "param1": "xx.xx.xx.xx",
	//
	//     "param2": "7d"
	//
	// }
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// The custom ID. If you do not specify this parameter when the playbook is triggered, a random ID is generated for fault locating and troubleshooting.
	//
	// example:
	//
	// f916b93e-e814-459f-9662-xxxxxxxxxx
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	// The task type. Valid values:
	//
	// 	- **command**
	//
	// 	- **playbook**
	//
	// example:
	//
	// playbook
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// example:
	//
	// f916b93e-e814-459f-9662-xxxxxxxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s TriggerSophonPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerSophonPlaybookRequest) GoString() string {
	return s.String()
}

func (s *TriggerSophonPlaybookRequest) SetCommandName(v string) *TriggerSophonPlaybookRequest {
	s.CommandName = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) SetInputParams(v string) *TriggerSophonPlaybookRequest {
	s.InputParams = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) SetSophonTaskId(v string) *TriggerSophonPlaybookRequest {
	s.SophonTaskId = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) SetTriggerType(v string) *TriggerSophonPlaybookRequest {
	s.TriggerType = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) SetUuid(v string) *TriggerSophonPlaybookRequest {
	s.Uuid = &v
	return s
}

type TriggerSophonPlaybookResponseBody struct {
	// The details that is returned after the command or playbook is triggered.
	Data *TriggerSophonPlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0DFC9403-54EB-5672-B690-9AA93C9EBB54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerSophonPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerSophonPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerSophonPlaybookResponseBody) SetData(v *TriggerSophonPlaybookResponseBodyData) *TriggerSophonPlaybookResponseBody {
	s.Data = v
	return s
}

func (s *TriggerSophonPlaybookResponseBody) SetRequestId(v string) *TriggerSophonPlaybookResponseBody {
	s.RequestId = &v
	return s
}

type TriggerSophonPlaybookResponseBodyData struct {
	// The custom ID. If you do not specify this parameter when the playbook is triggered, a random ID is generated for fault locating and troubleshooting.
	//
	// example:
	//
	// a7c6d055-a72f-4676-bc89-3cd9edc0284c
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
}

func (s TriggerSophonPlaybookResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TriggerSophonPlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *TriggerSophonPlaybookResponseBodyData) SetSophonTaskId(v string) *TriggerSophonPlaybookResponseBodyData {
	s.SophonTaskId = &v
	return s
}

type TriggerSophonPlaybookResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerSophonPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerSophonPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerSophonPlaybookResponse) GoString() string {
	return s.String()
}

func (s *TriggerSophonPlaybookResponse) SetHeaders(v map[string]*string) *TriggerSophonPlaybookResponse {
	s.Headers = v
	return s
}

func (s *TriggerSophonPlaybookResponse) SetStatusCode(v int32) *TriggerSophonPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerSophonPlaybookResponse) SetBody(v *TriggerSophonPlaybookResponseBody) *TriggerSophonPlaybookResponse {
	s.Body = v
	return s
}

type VerifyPlaybookRequest struct {
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9fcd3829-80ff-4681-be1e-4d2662c35fed
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The XML configuration of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// <?xml version="1.0" encoding="UTF-8"?><bpmn:definitions xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:bpmn="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:dc="http://www.omg.org/spec/DD/20100524/DC" targetNamespace="http://bpmn.io/schema/bpmn" id="Definitions_1"><bpmn:process id="Process_1" isExecutable="false"><bpmn:startEvent id="StartEvent_1"/></bpmn:process><bpmndi:BPMNDiagram id="BPMNDiagram_1"><bpmndi:BPMNPlane id="BPMNPlane_1" bpmnElement="Process_1"><bpmndi:BPMNShape id="_BPMNShape_StartEvent_2" bpmnElement="StartEvent_1"><dc:Bounds height="36.0" width="36.0" x="173.0" y="102.0"/></bpmndi:BPMNShape></bpmndi:BPMNPlane></bpmndi:BPMNDiagram></bpmn:definitions>
	TaskFlow *string `json:"TaskFlow,omitempty" xml:"TaskFlow,omitempty"`
}

func (s VerifyPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyPlaybookRequest) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookRequest) SetPlaybookUuid(v string) *VerifyPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *VerifyPlaybookRequest) SetTaskFlow(v string) *VerifyPlaybookRequest {
	s.TaskFlow = &v
	return s
}

type VerifyPlaybookResponseBody struct {
	// The result of the verification.
	CheckTaskInfos []*VerifyPlaybookResponseBodyCheckTaskInfos `json:"CheckTaskInfos,omitempty" xml:"CheckTaskInfos,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0DFC9403-54EB-5672-B690-9AA93C9EBB54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookResponseBody) SetCheckTaskInfos(v []*VerifyPlaybookResponseBodyCheckTaskInfos) *VerifyPlaybookResponseBody {
	s.CheckTaskInfos = v
	return s
}

func (s *VerifyPlaybookResponseBody) SetRequestId(v string) *VerifyPlaybookResponseBody {
	s.RequestId = &v
	return s
}

type VerifyPlaybookResponseBodyCheckTaskInfos struct {
	// The error message returned when the playbook does not pass the check.
	//
	// example:
	//
	// Node [python3_3] doesn\\"t have the asset information
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The name of the node in the playbook.
	//
	// example:
	//
	// python3_3
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The severity level of the verification information. Valid values:
	//
	// 	- warn: An issue may occur during playbook running.
	//
	// 	- error: The playbook cannot be compiled.
	//
	// 	- remind: The publishing and running of the playbook are not affected. We recommend that you optimize the playbook format.
	//
	// example:
	//
	// error
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s VerifyPlaybookResponseBodyCheckTaskInfos) String() string {
	return tea.Prettify(s)
}

func (s VerifyPlaybookResponseBodyCheckTaskInfos) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) SetDetail(v string) *VerifyPlaybookResponseBodyCheckTaskInfos {
	s.Detail = &v
	return s
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) SetNodeName(v string) *VerifyPlaybookResponseBodyCheckTaskInfos {
	s.NodeName = &v
	return s
}

func (s *VerifyPlaybookResponseBodyCheckTaskInfos) SetRiskLevel(v string) *VerifyPlaybookResponseBodyCheckTaskInfos {
	s.RiskLevel = &v
	return s
}

type VerifyPlaybookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyPlaybookResponse) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookResponse) SetHeaders(v map[string]*string) *VerifyPlaybookResponse {
	s.Headers = v
	return s
}

func (s *VerifyPlaybookResponse) SetStatusCode(v int32) *VerifyPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyPlaybookResponse) SetBody(v *VerifyPlaybookResponseBody) *VerifyPlaybookResponse {
	s.Body = v
	return s
}

type VerifyPythonFileRequest struct {
	// The Python code snippet.
	//
	// This parameter is required.
	//
	// example:
	//
	// import logging
	//
	// def execute (params):
	//
	//   success=True
	//
	//   message=\\"OK\\"
	//
	//   data=[]
	//
	//   return (success,message,data)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s VerifyPythonFileRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyPythonFileRequest) GoString() string {
	return s.String()
}

func (s *VerifyPythonFileRequest) SetContent(v string) *VerifyPythonFileRequest {
	s.Content = &v
	return s
}

type VerifyPythonFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F72685FB-A6E6-5A9A-97F7-6DC1056E63CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The verification result. If the parameter is left empty, the syntax of the code snippet is correct.
	Syntax []*VerifyPythonFileResponseBodySyntax `json:"Syntax,omitempty" xml:"Syntax,omitempty" type:"Repeated"`
}

func (s VerifyPythonFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyPythonFileResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyPythonFileResponseBody) SetRequestId(v string) *VerifyPythonFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyPythonFileResponseBody) SetSyntax(v []*VerifyPythonFileResponseBodySyntax) *VerifyPythonFileResponseBody {
	s.Syntax = v
	return s
}

type VerifyPythonFileResponseBodySyntax struct {
	// The number that indicates the end column of the error code.
	//
	// example:
	//
	// 5
	EndColumn *int32 `json:"EndColumn,omitempty" xml:"EndColumn,omitempty"`
	// The number that indicates the end line of the error code.
	//
	// example:
	//
	// 5
	EndLineNumber *int32 `json:"EndLineNumber,omitempty" xml:"EndLineNumber,omitempty"`
	// The error message for the error code.
	//
	// example:
	//
	// undefined name \\"ab\\"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The severity level of the error code. Valid values:
	//
	// 	- 4: moderate
	//
	// 	- 8: serious
	//
	// example:
	//
	// 4
	Severity *int32 `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The number that indicates the start column of the error code.
	//
	// example:
	//
	// 2
	StartColumn *int32 `json:"StartColumn,omitempty" xml:"StartColumn,omitempty"`
	// The number that indicates the start line of the error code.
	//
	// example:
	//
	// 2
	StartLineNumber *int32 `json:"StartLineNumber,omitempty" xml:"StartLineNumber,omitempty"`
}

func (s VerifyPythonFileResponseBodySyntax) String() string {
	return tea.Prettify(s)
}

func (s VerifyPythonFileResponseBodySyntax) GoString() string {
	return s.String()
}

func (s *VerifyPythonFileResponseBodySyntax) SetEndColumn(v int32) *VerifyPythonFileResponseBodySyntax {
	s.EndColumn = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetEndLineNumber(v int32) *VerifyPythonFileResponseBodySyntax {
	s.EndLineNumber = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetMessage(v string) *VerifyPythonFileResponseBodySyntax {
	s.Message = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetSeverity(v int32) *VerifyPythonFileResponseBodySyntax {
	s.Severity = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetStartColumn(v int32) *VerifyPythonFileResponseBodySyntax {
	s.StartColumn = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetStartLineNumber(v int32) *VerifyPythonFileResponseBodySyntax {
	s.StartLineNumber = &v
	return s
}

type VerifyPythonFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyPythonFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyPythonFileResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyPythonFileResponse) GoString() string {
	return s.String()
}

func (s *VerifyPythonFileResponse) SetHeaders(v map[string]*string) *VerifyPythonFileResponse {
	s.Headers = v
	return s
}

func (s *VerifyPythonFileResponse) SetStatusCode(v int32) *VerifyPythonFileResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyPythonFileResponse) SetBody(v *VerifyPythonFileResponseBody) *VerifyPythonFileResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sophonsoar"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Modifies the statuses of playbooks at a time.
//
// @param request - BatchModifyInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchModifyInstanceStatusResponse
func (client *Client) BatchModifyInstanceStatusWithOptions(request *BatchModifyInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *BatchModifyInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["Active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchModifyInstanceStatus"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &BatchModifyInstanceStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &BatchModifyInstanceStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the statuses of playbooks at a time.
//
// @param request - BatchModifyInstanceStatusRequest
//
// @return BatchModifyInstanceStatusResponse
func (client *Client) BatchModifyInstanceStatus(request *BatchModifyInstanceStatusRequest) (_result *BatchModifyInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchModifyInstanceStatusResponse{}
	_body, _err := client.BatchModifyInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Compares configurations between two versions of a published playbook.
//
// @param request - ComparePlaybooksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ComparePlaybooksResponse
func (client *Client) ComparePlaybooksWithOptions(request *ComparePlaybooksRequest, runtime *util.RuntimeOptions) (_result *ComparePlaybooksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NewPlaybookReleaseId)) {
		query["NewPlaybookReleaseId"] = request.NewPlaybookReleaseId
	}

	if !tea.BoolValue(util.IsUnset(request.OldPlaybookReleaseId)) {
		query["OldPlaybookReleaseId"] = request.OldPlaybookReleaseId
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		query["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ComparePlaybooks"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ComparePlaybooksResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ComparePlaybooksResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Compares configurations between two versions of a published playbook.
//
// @param request - ComparePlaybooksRequest
//
// @return ComparePlaybooksResponse
func (client *Client) ComparePlaybooks(request *ComparePlaybooksRequest) (_result *ComparePlaybooksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ComparePlaybooksResponse{}
	_body, _err := client.ComparePlaybooksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Convert XML configuration.
//
// Description:
//
// Please ensure that you fully understand the billing method and [pricing](https://www.aliyun.com/price/product#/sas/detail/sas) of the orchestration product before using this interface.
//
// @param request - ConvertPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertPlaybookResponse
func (client *Client) ConvertPlaybookWithOptions(request *ConvertPlaybookRequest, runtime *util.RuntimeOptions) (_result *ConvertPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		query["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		query["RoleType"] = request.RoleType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Taskflow)) {
		body["Taskflow"] = request.Taskflow
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConvertPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ConvertPlaybookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ConvertPlaybookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Convert XML configuration.
//
// Description:
//
// Please ensure that you fully understand the billing method and [pricing](https://www.aliyun.com/price/product#/sas/detail/sas) of the orchestration product before using this interface.
//
// @param request - ConvertPlaybookRequest
//
// @return ConvertPlaybookResponse
func (client *Client) ConvertPlaybook(request *ConvertPlaybookRequest) (_result *ConvertPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertPlaybookResponse{}
	_body, _err := client.ConvertPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 剧本复制
//
// @param request - CopyPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyPlaybookResponse
func (client *Client) CopyPlaybookWithOptions(request *CopyPlaybookRequest, runtime *util.RuntimeOptions) (_result *CopyPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		query["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		query["RoleType"] = request.RoleType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseVersion)) {
		body["ReleaseVersion"] = request.ReleaseVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePlaybookUuid)) {
		body["SourcePlaybookUuid"] = request.SourcePlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CopyPlaybookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CopyPlaybookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 剧本复制
//
// @param request - CopyPlaybookRequest
//
// @return CopyPlaybookResponse
func (client *Client) CopyPlaybook(request *CopyPlaybookRequest) (_result *CopyPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyPlaybookResponse{}
	_body, _err := client.CopyPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// New Playbook.
//
// Description:
//
// Create Playbook.
//
// @param request - CreatePlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePlaybookResponse
func (client *Client) CreatePlaybookWithOptions(request *CreatePlaybookRequest, runtime *util.RuntimeOptions) (_result *CreatePlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		body["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.TaskflowType)) {
		body["TaskflowType"] = request.TaskflowType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePlaybookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePlaybookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// New Playbook.
//
// Description:
//
// Create Playbook.
//
// @param request - CreatePlaybookRequest
//
// @return CreatePlaybookResponse
func (client *Client) CreatePlaybook(request *CreatePlaybookRequest) (_result *CreatePlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePlaybookResponse{}
	_body, _err := client.CreatePlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Debugs a playbook.
//
// @param request - DebugPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugPlaybookResponse
func (client *Client) DebugPlaybookWithOptions(request *DebugPlaybookRequest, runtime *util.RuntimeOptions) (_result *DebugPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		body["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Record)) {
		body["Record"] = request.Record
	}

	if !tea.BoolValue(util.IsUnset(request.Taskflow)) {
		body["Taskflow"] = request.Taskflow
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DebugPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DebugPlaybookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DebugPlaybookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Debugs a playbook.
//
// @param request - DebugPlaybookRequest
//
// @return DebugPlaybookResponse
func (client *Client) DebugPlaybook(request *DebugPlaybookRequest) (_result *DebugPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DebugPlaybookResponse{}
	_body, _err := client.DebugPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the assets in a component.
//
// @param request - DeleteComponentAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComponentAssetResponse
func (client *Client) DeleteComponentAssetWithOptions(request *DeleteComponentAssetRequest, runtime *util.RuntimeOptions) (_result *DeleteComponentAssetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetId)) {
		query["AssetId"] = request.AssetId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteComponentAsset"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteComponentAssetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteComponentAssetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes the assets in a component.
//
// @param request - DeleteComponentAssetRequest
//
// @return DeleteComponentAssetResponse
func (client *Client) DeleteComponentAsset(request *DeleteComponentAssetRequest) (_result *DeleteComponentAssetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteComponentAssetResponse{}
	_body, _err := client.DeleteComponentAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom playbook.
//
// @param request - DeletePlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePlaybookResponse
func (client *Client) DeletePlaybookWithOptions(request *DeletePlaybookRequest, runtime *util.RuntimeOptions) (_result *DeletePlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		body["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeletePlaybookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeletePlaybookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a custom playbook.
//
// @param request - DeletePlaybookRequest
//
// @return DeletePlaybookResponse
func (client *Client) DeletePlaybook(request *DeletePlaybookRequest) (_result *DeletePlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePlaybookResponse{}
	_body, _err := client.DeletePlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metadata of assets in a component. The metadata of an asset refers to the fields that describe the asset.
//
// @param request - DescribeComponentAssetFormRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentAssetFormResponse
func (client *Client) DescribeComponentAssetFormWithOptions(request *DescribeComponentAssetFormRequest, runtime *util.RuntimeOptions) (_result *DescribeComponentAssetFormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponentAssetForm"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeComponentAssetFormResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeComponentAssetFormResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the metadata of assets in a component. The metadata of an asset refers to the fields that describe the asset.
//
// @param request - DescribeComponentAssetFormRequest
//
// @return DescribeComponentAssetFormResponse
func (client *Client) DescribeComponentAssetForm(request *DescribeComponentAssetFormRequest) (_result *DescribeComponentAssetFormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComponentAssetFormResponse{}
	_body, _err := client.DescribeComponentAssetFormWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of assets in a component.
//
// @param request - DescribeComponentAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentAssetsResponse
func (client *Client) DescribeComponentAssetsWithOptions(request *DescribeComponentAssetsRequest, runtime *util.RuntimeOptions) (_result *DescribeComponentAssetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponentAssets"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeComponentAssetsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeComponentAssetsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of assets in a component.
//
// @param request - DescribeComponentAssetsRequest
//
// @return DescribeComponentAssetsResponse
func (client *Client) DescribeComponentAssets(request *DescribeComponentAssetsRequest) (_result *DescribeComponentAssetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComponentAssetsResponse{}
	_body, _err := client.DescribeComponentAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of common components that are available.
//
// @param request - DescribeComponentListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentListResponse
func (client *Client) DescribeComponentListWithOptions(request *DescribeComponentListRequest, runtime *util.RuntimeOptions) (_result *DescribeComponentListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponentList"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeComponentListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeComponentListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of common components that are available.
//
// @param request - DescribeComponentListRequest
//
// @return DescribeComponentListResponse
func (client *Client) DescribeComponentList(request *DescribeComponentListRequest) (_result *DescribeComponentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComponentListResponse{}
	_body, _err := client.DescribeComponentListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of predefined components that are available.
//
// @param request - DescribeComponentPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentPlaybookResponse
func (client *Client) DescribeComponentPlaybookWithOptions(request *DescribeComponentPlaybookRequest, runtime *util.RuntimeOptions) (_result *DescribeComponentPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponentPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeComponentPlaybookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeComponentPlaybookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of predefined components that are available.
//
// @param request - DescribeComponentPlaybookRequest
//
// @return DescribeComponentPlaybookResponse
func (client *Client) DescribeComponentPlaybook(request *DescribeComponentPlaybookRequest) (_result *DescribeComponentPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComponentPlaybookResponse{}
	_body, _err := client.DescribeComponentPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the JavaScript file of a component. The component uses the returned JavaScript file for page rendering.
//
// @param request - DescribeComponentsJsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentsJsResponse
func (client *Client) DescribeComponentsJsWithOptions(request *DescribeComponentsJsRequest, runtime *util.RuntimeOptions) (_result *DescribeComponentsJsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponentsJs"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeComponentsJsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeComponentsJsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the JavaScript file of a component. The component uses the returned JavaScript file for page rendering.
//
// @param request - DescribeComponentsJsRequest
//
// @return DescribeComponentsJsResponse
func (client *Client) DescribeComponentsJs(request *DescribeComponentsJsRequest) (_result *DescribeComponentsJsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComponentsJsResponse{}
	_body, _err := client.DescribeComponentsJsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the published versions of a playbook after deduplication.
//
// @param request - DescribeDistinctReleasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDistinctReleasesResponse
func (client *Client) DescribeDistinctReleasesWithOptions(request *DescribeDistinctReleasesRequest, runtime *util.RuntimeOptions) (_result *DescribeDistinctReleasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDistinctReleases"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDistinctReleasesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDistinctReleasesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about the published versions of a playbook after deduplication.
//
// @param request - DescribeDistinctReleasesRequest
//
// @return DescribeDistinctReleasesResponse
func (client *Client) DescribeDistinctReleases(request *DescribeDistinctReleasesRequest) (_result *DescribeDistinctReleasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDistinctReleasesResponse{}
	_body, _err := client.DescribeDistinctReleasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries enumeration items that are required by a cloud service.
//
// @param request - DescribeEnumItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnumItemsResponse
func (client *Client) DescribeEnumItemsWithOptions(request *DescribeEnumItemsRequest, runtime *util.RuntimeOptions) (_result *DescribeEnumItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEnumItems"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeEnumItemsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeEnumItemsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries enumeration items that are required by a cloud service.
//
// @param request - DescribeEnumItemsRequest
//
// @return DescribeEnumItemsResponse
func (client *Client) DescribeEnumItems(request *DescribeEnumItemsRequest) (_result *DescribeEnumItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEnumItemsResponse{}
	_body, _err := client.DescribeEnumItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the playbooks that are available for an automatic response plan.
//
// @param request - DescribeExecutePlaybooksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExecutePlaybooksResponse
func (client *Client) DescribeExecutePlaybooksWithOptions(request *DescribeExecutePlaybooksRequest, runtime *util.RuntimeOptions) (_result *DescribeExecutePlaybooksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExecutePlaybooks"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeExecutePlaybooksResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeExecutePlaybooksResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the playbooks that are available for an automatic response plan.
//
// @param request - DescribeExecutePlaybooksRequest
//
// @return DescribeExecutePlaybooksResponse
func (client *Client) DescribeExecutePlaybooks(request *DescribeExecutePlaybooksRequest) (_result *DescribeExecutePlaybooksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExecutePlaybooksResponse{}
	_body, _err := client.DescribeExecutePlaybooksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the global configuration information about a cloud service.
//
// @param request - DescribeFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFieldResponse
func (client *Client) DescribeFieldWithOptions(request *DescribeFieldRequest, runtime *util.RuntimeOptions) (_result *DescribeFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeField"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeFieldResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeFieldResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the global configuration information about a cloud service.
//
// @param request - DescribeFieldRequest
//
// @return DescribeFieldResponse
func (client *Client) DescribeField(request *DescribeFieldRequest) (_result *DescribeFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFieldResponse{}
	_body, _err := client.DescribeFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取OpenAPI的产品列表
//
// @param request - DescribeGroupProductionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupProductionsResponse
func (client *Client) DescribeGroupProductionsWithOptions(request *DescribeGroupProductionsRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupProductionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGroupProductions"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeGroupProductionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeGroupProductionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取OpenAPI的产品列表
//
// @param request - DescribeGroupProductionsRequest
//
// @return DescribeGroupProductionsResponse
func (client *Client) DescribeGroupProductions(request *DescribeGroupProductionsRequest) (_result *DescribeGroupProductionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupProductionsResponse{}
	_body, _err := client.DescribeGroupProductionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the output structure information of each node in a playbook based on the most recent running record of the playbook.
//
// @param request - DescribeLatestRecordSchemaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLatestRecordSchemaResponse
func (client *Client) DescribeLatestRecordSchemaWithOptions(request *DescribeLatestRecordSchemaRequest, runtime *util.RuntimeOptions) (_result *DescribeLatestRecordSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLatestRecordSchema"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeLatestRecordSchemaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeLatestRecordSchemaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the output structure information of each node in a playbook based on the most recent running record of the playbook.
//
// @param request - DescribeLatestRecordSchemaRequest
//
// @return DescribeLatestRecordSchemaResponse
func (client *Client) DescribeLatestRecordSchema(request *DescribeLatestRecordSchemaRequest) (_result *DescribeLatestRecordSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLatestRecordSchemaResponse{}
	_body, _err := client.DescribeLatestRecordSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries recommended dynamic input parameters of a component for playbook orchestration.
//
// @param request - DescribeNodeParamTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodeParamTagsResponse
func (client *Client) DescribeNodeParamTagsWithOptions(request *DescribeNodeParamTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeParamTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNodeParamTags"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeNodeParamTagsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeNodeParamTagsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries recommended dynamic input parameters of a component for playbook orchestration.
//
// @param request - DescribeNodeParamTagsRequest
//
// @return DescribeNodeParamTagsResponse
func (client *Client) DescribeNodeParamTags(request *DescribeNodeParamTagsRequest) (_result *DescribeNodeParamTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeParamTagsResponse{}
	_body, _err := client.DescribeNodeParamTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the nodes that reference the same node in a playbook.
//
// @param request - DescribeNodeUsedInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodeUsedInfosResponse
func (client *Client) DescribeNodeUsedInfosWithOptions(request *DescribeNodeUsedInfosRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeUsedInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNodeUsedInfos"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeNodeUsedInfosResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeNodeUsedInfosResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the nodes that reference the same node in a playbook.
//
// @param request - DescribeNodeUsedInfosRequest
//
// @return DescribeNodeUsedInfosResponse
func (client *Client) DescribeNodeUsedInfos(request *DescribeNodeUsedInfosRequest) (_result *DescribeNodeUsedInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeUsedInfosResponse{}
	_body, _err := client.DescribeNodeUsedInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询通知消息模版列表
//
// @param request - DescribeNotifyTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNotifyTemplateListResponse
func (client *Client) DescribeNotifyTemplateListWithOptions(request *DescribeNotifyTemplateListRequest, runtime *util.RuntimeOptions) (_result *DescribeNotifyTemplateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNotifyTemplateList"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeNotifyTemplateListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeNotifyTemplateListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询通知消息模版列表
//
// @param request - DescribeNotifyTemplateListRequest
//
// @return DescribeNotifyTemplateListResponse
func (client *Client) DescribeNotifyTemplateList(request *DescribeNotifyTemplateListRequest) (_result *DescribeNotifyTemplateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNotifyTemplateListResponse{}
	_body, _err := client.DescribeNotifyTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取产品接口的详情
//
// @param request - DescribeOpenApiInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpenApiInfoResponse
func (client *Client) DescribeOpenApiInfoWithOptions(request *DescribeOpenApiInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeOpenApiInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOpenApiInfo"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeOpenApiInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeOpenApiInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取产品接口的详情
//
// @param request - DescribeOpenApiInfoRequest
//
// @return DescribeOpenApiInfoResponse
func (client *Client) DescribeOpenApiInfo(request *DescribeOpenApiInfoRequest) (_result *DescribeOpenApiInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpenApiInfoResponse{}
	_body, _err := client.DescribeOpenApiInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取产品的接口列表
//
// @param request - DescribeOpenApiListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpenApiListResponse
func (client *Client) DescribeOpenApiListWithOptions(request *DescribeOpenApiListRequest, runtime *util.RuntimeOptions) (_result *DescribeOpenApiListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOpenApiList"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeOpenApiListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeOpenApiListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取产品的接口列表
//
// @param request - DescribeOpenApiListRequest
//
// @return DescribeOpenApiListResponse
func (client *Client) DescribeOpenApiList(request *DescribeOpenApiListRequest) (_result *DescribeOpenApiListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpenApiListResponse{}
	_body, _err := client.DescribeOpenApiListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the XML configuration of a playbook.
//
// @param request - DescribePlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybookResponse
func (client *Client) DescribePlaybookWithOptions(request *DescribePlaybookRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribePlaybookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribePlaybookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the XML configuration of a playbook.
//
// @param request - DescribePlaybookRequest
//
// @return DescribePlaybookResponse
func (client *Client) DescribePlaybook(request *DescribePlaybookRequest) (_result *DescribePlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybookResponse{}
	_body, _err := client.DescribePlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the input and output parameter configurations of a playbook.
//
// @param request - DescribePlaybookInputOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybookInputOutputResponse
func (client *Client) DescribePlaybookInputOutputWithOptions(request *DescribePlaybookInputOutputRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybookInputOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybookInputOutput"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribePlaybookInputOutputResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribePlaybookInputOutputResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the input and output parameter configurations of a playbook.
//
// @param request - DescribePlaybookInputOutputRequest
//
// @return DescribePlaybookInputOutputResponse
func (client *Client) DescribePlaybookInputOutput(request *DescribePlaybookInputOutputRequest) (_result *DescribePlaybookInputOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybookInputOutputResponse{}
	_body, _err := client.DescribePlaybookInputOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metrics of a playbook. The metrics include the playbook name, playbook description, the number of times that the playbook is run, and the failure rate of the playbook.
//
// @param request - DescribePlaybookMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybookMetricsResponse
func (client *Client) DescribePlaybookMetricsWithOptions(request *DescribePlaybookMetricsRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybookMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybookMetrics"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribePlaybookMetricsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribePlaybookMetricsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the metrics of a playbook. The metrics include the playbook name, playbook description, the number of times that the playbook is run, and the failure rate of the playbook.
//
// @param request - DescribePlaybookMetricsRequest
//
// @return DescribePlaybookMetricsResponse
func (client *Client) DescribePlaybookMetrics(request *DescribePlaybookMetricsRequest) (_result *DescribePlaybookMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybookMetricsResponse{}
	_body, _err := client.DescribePlaybookMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical output data of a component node.
//
// @param request - DescribePlaybookNodesOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybookNodesOutputResponse
func (client *Client) DescribePlaybookNodesOutputWithOptions(request *DescribePlaybookNodesOutputRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybookNodesOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybookNodesOutput"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribePlaybookNodesOutputResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribePlaybookNodesOutputResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the historical output data of a component node.
//
// @param request - DescribePlaybookNodesOutputRequest
//
// @return DescribePlaybookNodesOutputResponse
func (client *Client) DescribePlaybookNodesOutput(request *DescribePlaybookNodesOutputRequest) (_result *DescribePlaybookNodesOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybookNodesOutputResponse{}
	_body, _err := client.DescribePlaybookNodesOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics of Security Orchestration Automation Response (SOAR), such as the numbers of created and enabled playbooks.
//
// @param request - DescribePlaybookNumberMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybookNumberMetricsResponse
func (client *Client) DescribePlaybookNumberMetricsWithOptions(request *DescribePlaybookNumberMetricsRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybookNumberMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybookNumberMetrics"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribePlaybookNumberMetricsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribePlaybookNumberMetricsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the statistics of Security Orchestration Automation Response (SOAR), such as the numbers of created and enabled playbooks.
//
// @param request - DescribePlaybookNumberMetricsRequest
//
// @return DescribePlaybookNumberMetricsResponse
func (client *Client) DescribePlaybookNumberMetrics(request *DescribePlaybookNumberMetricsRequest) (_result *DescribePlaybookNumberMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybookNumberMetricsResponse{}
	_body, _err := client.DescribePlaybookNumberMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the published versions of a playbook.
//
// @param request - DescribePlaybookReleasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybookReleasesResponse
func (client *Client) DescribePlaybookReleasesWithOptions(request *DescribePlaybookReleasesRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybookReleasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybookReleases"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribePlaybookReleasesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribePlaybookReleasesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about the published versions of a playbook.
//
// @param request - DescribePlaybookReleasesRequest
//
// @return DescribePlaybookReleasesResponse
func (client *Client) DescribePlaybookReleases(request *DescribePlaybookReleasesRequest) (_result *DescribePlaybookReleasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybookReleasesResponse{}
	_body, _err := client.DescribePlaybookReleasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve the list of playbooks.
//
// @param request - DescribePlaybooksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlaybooksResponse
func (client *Client) DescribePlaybooksWithOptions(request *DescribePlaybooksRequest, runtime *util.RuntimeOptions) (_result *DescribePlaybooksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlaybooks"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribePlaybooksResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribePlaybooksResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Retrieve the list of playbooks.
//
// @param request - DescribePlaybooksRequest
//
// @return DescribePlaybooksResponse
func (client *Client) DescribePlaybooks(request *DescribePlaybooksRequest) (_result *DescribePlaybooksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlaybooksResponse{}
	_body, _err := client.DescribePlaybooksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an API operation.
//
// @param request - DescribePopApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePopApiResponse
func (client *Client) DescribePopApiWithOptions(request *DescribePopApiRequest, runtime *util.RuntimeOptions) (_result *DescribePopApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePopApi"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribePopApiResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribePopApiResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of an API operation.
//
// @param request - DescribePopApiRequest
//
// @return DescribePopApiResponse
func (client *Client) DescribePopApi(request *DescribePopApiRequest) (_result *DescribePopApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePopApiResponse{}
	_body, _err := client.DescribePopApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of API operations for an Alibaba Cloud service.
//
// @param request - DescribePopApiItemListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePopApiItemListResponse
func (client *Client) DescribePopApiItemListWithOptions(request *DescribePopApiItemListRequest, runtime *util.RuntimeOptions) (_result *DescribePopApiItemListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePopApiItemList"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribePopApiItemListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribePopApiItemListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of API operations for an Alibaba Cloud service.
//
// @param request - DescribePopApiItemListRequest
//
// @return DescribePopApiItemListResponse
func (client *Client) DescribePopApiItemList(request *DescribePopApiItemListRequest) (_result *DescribePopApiItemListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePopApiItemListResponse{}
	_body, _err := client.DescribePopApiItemListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取统计信息
//
// @param request - DescribeProcessStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProcessStatisticsResponse
func (client *Client) DescribeProcessStatisticsWithOptions(request *DescribeProcessStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeProcessStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProcessStatistics"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeProcessStatisticsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeProcessStatisticsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取统计信息
//
// @param request - DescribeProcessStatisticsRequest
//
// @return DescribeProcessStatisticsResponse
func (client *Client) DescribeProcessStatistics(request *DescribeProcessStatisticsRequest) (_result *DescribeProcessStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProcessStatisticsResponse{}
	_body, _err := client.DescribeProcessStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the number of associated disposal tasks based on the entity UUID.
//
// @param request - DescribeProcessTaskCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProcessTaskCountResponse
func (client *Client) DescribeProcessTaskCountWithOptions(request *DescribeProcessTaskCountRequest, runtime *util.RuntimeOptions) (_result *DescribeProcessTaskCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProcessTaskCount"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeProcessTaskCountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeProcessTaskCountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Query the number of associated disposal tasks based on the entity UUID.
//
// @param request - DescribeProcessTaskCountRequest
//
// @return DescribeProcessTaskCountResponse
func (client *Client) DescribeProcessTaskCount(request *DescribeProcessTaskCountRequest) (_result *DescribeProcessTaskCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProcessTaskCountResponse{}
	_body, _err := client.DescribeProcessTaskCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about handling tasks. When you use Security Orchestration Automation Response (SOAR) to handle events, handling tasks are generated in the handling center.
//
// @param request - DescribeProcessTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProcessTasksResponse
func (client *Client) DescribeProcessTasksWithOptions(request *DescribeProcessTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeProcessTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		query["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.EntityUuid)) {
		query["EntityUuid"] = request.EntityUuid
	}

	if !tea.BoolValue(util.IsUnset(request.EventUuid)) {
		query["EventUuid"] = request.EventUuid
	}

	if !tea.BoolValue(util.IsUnset(request.OrderField)) {
		query["OrderField"] = request.OrderField
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParamContent)) {
		query["ParamContent"] = request.ParamContent
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessActionEnd)) {
		query["ProcessActionEnd"] = request.ProcessActionEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessActionStart)) {
		query["ProcessActionStart"] = request.ProcessActionStart
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessRemoveEnd)) {
		query["ProcessRemoveEnd"] = request.ProcessRemoveEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessRemoveStart)) {
		query["ProcessRemoveStart"] = request.ProcessRemoveStart
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessStrategyUuid)) {
		query["ProcessStrategyUuid"] = request.ProcessStrategyUuid
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		query["TaskStatus"] = request.TaskStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerSource)) {
		query["TriggerSource"] = request.TriggerSource
	}

	if !tea.BoolValue(util.IsUnset(request.YunCode)) {
		query["YunCode"] = request.YunCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProcessTasks"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeProcessTasksResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeProcessTasksResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about handling tasks. When you use Security Orchestration Automation Response (SOAR) to handle events, handling tasks are generated in the handling center.
//
// @param request - DescribeProcessTasksRequest
//
// @return DescribeProcessTasksResponse
func (client *Client) DescribeProcessTasks(request *DescribeProcessTasksRequest) (_result *DescribeProcessTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProcessTasksResponse{}
	_body, _err := client.DescribeProcessTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data that is returned when a component initiates an action in a playbook task.
//
// @param request - DescribeSoarRecordActionOutputListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSoarRecordActionOutputListResponse
func (client *Client) DescribeSoarRecordActionOutputListWithOptions(request *DescribeSoarRecordActionOutputListRequest, runtime *util.RuntimeOptions) (_result *DescribeSoarRecordActionOutputListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSoarRecordActionOutputList"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSoarRecordActionOutputListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSoarRecordActionOutputListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the data that is returned when a component initiates an action in a playbook task.
//
// @param request - DescribeSoarRecordActionOutputListRequest
//
// @return DescribeSoarRecordActionOutputListResponse
func (client *Client) DescribeSoarRecordActionOutputList(request *DescribeSoarRecordActionOutputListRequest) (_result *DescribeSoarRecordActionOutputListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSoarRecordActionOutputListResponse{}
	_body, _err := client.DescribeSoarRecordActionOutputListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the input and output data of a component action. You can call this operation after a playbook is run.
//
// @param request - DescribeSoarRecordInOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSoarRecordInOutputResponse
func (client *Client) DescribeSoarRecordInOutputWithOptions(request *DescribeSoarRecordInOutputRequest, runtime *util.RuntimeOptions) (_result *DescribeSoarRecordInOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSoarRecordInOutput"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSoarRecordInOutputResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSoarRecordInOutputResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the input and output data of a component action. You can call this operation after a playbook is run.
//
// @param request - DescribeSoarRecordInOutputRequest
//
// @return DescribeSoarRecordInOutputResponse
func (client *Client) DescribeSoarRecordInOutput(request *DescribeSoarRecordInOutputRequest) (_result *DescribeSoarRecordInOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSoarRecordInOutputResponse{}
	_body, _err := client.DescribeSoarRecordInOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get the execution records of a playbook.
//
// @param request - DescribeSoarRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSoarRecordsResponse
func (client *Client) DescribeSoarRecordsWithOptions(request *DescribeSoarRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSoarRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSoarRecords"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSoarRecordsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSoarRecordsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Get the execution records of a playbook.
//
// @param request - DescribeSoarRecordsRequest
//
// @return DescribeSoarRecordsResponse
func (client *Client) DescribeSoarRecords(request *DescribeSoarRecordsRequest) (_result *DescribeSoarRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSoarRecordsResponse{}
	_body, _err := client.DescribeSoarRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution records of a component during the running of a playbook.
//
// @param request - DescribeSoarTaskAndActionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSoarTaskAndActionsResponse
func (client *Client) DescribeSoarTaskAndActionsWithOptions(request *DescribeSoarTaskAndActionsRequest, runtime *util.RuntimeOptions) (_result *DescribeSoarTaskAndActionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSoarTaskAndActions"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSoarTaskAndActionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSoarTaskAndActionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the execution records of a component during the running of a playbook.
//
// @param request - DescribeSoarTaskAndActionsRequest
//
// @return DescribeSoarTaskAndActionsResponse
func (client *Client) DescribeSoarTaskAndActions(request *DescribeSoarTaskAndActionsRequest) (_result *DescribeSoarTaskAndActionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSoarTaskAndActionsResponse{}
	_body, _err := client.DescribeSoarTaskAndActionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the commands that can be run to obtain objects.
//
// @param request - DescribeSophonCommandsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSophonCommandsResponse
func (client *Client) DescribeSophonCommandsWithOptions(request *DescribeSophonCommandsRequest, runtime *util.RuntimeOptions) (_result *DescribeSophonCommandsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSophonCommands"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSophonCommandsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSophonCommandsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the commands that can be run to obtain objects.
//
// @param request - DescribeSophonCommandsRequest
//
// @return DescribeSophonCommandsResponse
func (client *Client) DescribeSophonCommands(request *DescribeSophonCommandsRequest) (_result *DescribeSophonCommandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSophonCommandsResponse{}
	_body, _err := client.DescribeSophonCommandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询云厂商OpenApi列表
//
// @param request - DescribeVendorApiListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVendorApiListResponse
func (client *Client) DescribeVendorApiListWithOptions(request *DescribeVendorApiListRequest, runtime *util.RuntimeOptions) (_result *DescribeVendorApiListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.VendorCode)) {
		query["VendorCode"] = request.VendorCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVendorApiList"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeVendorApiListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeVendorApiListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询云厂商OpenApi列表
//
// @param request - DescribeVendorApiListRequest
//
// @return DescribeVendorApiListResponse
func (client *Client) DescribeVendorApiList(request *DescribeVendorApiListRequest) (_result *DescribeVendorApiListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVendorApiListResponse{}
	_body, _err := client.DescribeVendorApiListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operational logs of a Python3 script by using the UUID that is returned when the script is run. The UUID is specified by requestUuid.
//
// @param request - DescriberPython3ScriptLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescriberPython3ScriptLogsResponse
func (client *Client) DescriberPython3ScriptLogsWithOptions(request *DescriberPython3ScriptLogsRequest, runtime *util.RuntimeOptions) (_result *DescriberPython3ScriptLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescriberPython3ScriptLogs"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescriberPython3ScriptLogsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescriberPython3ScriptLogsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the operational logs of a Python3 script by using the UUID that is returned when the script is run. The UUID is specified by requestUuid.
//
// @param request - DescriberPython3ScriptLogsRequest
//
// @return DescriberPython3ScriptLogsResponse
func (client *Client) DescriberPython3ScriptLogs(request *DescriberPython3ScriptLogsRequest) (_result *DescriberPython3ScriptLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescriberPython3ScriptLogsResponse{}
	_body, _err := client.DescriberPython3ScriptLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about the asset that is configured for a component.
//
// @param request - ModifyComponentAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyComponentAssetResponse
func (client *Client) ModifyComponentAssetWithOptions(request *ModifyComponentAssetRequest, runtime *util.RuntimeOptions) (_result *ModifyComponentAssetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetConfig)) {
		query["AssetConfig"] = request.AssetConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyComponentAsset"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyComponentAssetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyComponentAssetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the information about the asset that is configured for a component.
//
// @param request - ModifyComponentAssetRequest
//
// @return ModifyComponentAssetResponse
func (client *Client) ModifyComponentAsset(request *ModifyComponentAssetRequest) (_result *ModifyComponentAssetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyComponentAssetResponse{}
	_body, _err := client.ModifyComponentAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a playbook.
//
// @param request - ModifyPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPlaybookResponse
func (client *Client) ModifyPlaybookWithOptions(request *ModifyPlaybookRequest, runtime *util.RuntimeOptions) (_result *ModifyPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		body["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Taskflow)) {
		body["Taskflow"] = request.Taskflow
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyPlaybookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyPlaybookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the configuration of a playbook.
//
// @param request - ModifyPlaybookRequest
//
// @return ModifyPlaybookResponse
func (client *Client) ModifyPlaybook(request *ModifyPlaybookRequest) (_result *ModifyPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPlaybookResponse{}
	_body, _err := client.ModifyPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the input and output parameters of a playbook.
//
// @param request - ModifyPlaybookInputOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPlaybookInputOutputResponse
func (client *Client) ModifyPlaybookInputOutputWithOptions(request *ModifyPlaybookInputOutputRequest, runtime *util.RuntimeOptions) (_result *ModifyPlaybookInputOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExeConfig)) {
		body["ExeConfig"] = request.ExeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.InputParams)) {
		body["InputParams"] = request.InputParams
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		body["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OutputParams)) {
		body["OutputParams"] = request.OutputParams
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		body["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPlaybookInputOutput"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyPlaybookInputOutputResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyPlaybookInputOutputResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the input and output parameters of a playbook.
//
// @param request - ModifyPlaybookInputOutputRequest
//
// @return ModifyPlaybookInputOutputResponse
func (client *Client) ModifyPlaybookInputOutput(request *ModifyPlaybookInputOutputRequest) (_result *ModifyPlaybookInputOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPlaybookInputOutputResponse{}
	_body, _err := client.ModifyPlaybookInputOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the status of a playbook.
//
// @param request - ModifyPlaybookInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPlaybookInstanceStatusResponse
func (client *Client) ModifyPlaybookInstanceStatusWithOptions(request *ModifyPlaybookInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyPlaybookInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["Active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPlaybookInstanceStatus"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyPlaybookInstanceStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyPlaybookInstanceStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the status of a playbook.
//
// @param request - ModifyPlaybookInstanceStatusRequest
//
// @return ModifyPlaybookInstanceStatusResponse
func (client *Client) ModifyPlaybookInstanceStatus(request *ModifyPlaybookInstanceStatusRequest) (_result *ModifyPlaybookInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPlaybookInstanceStatusResponse{}
	_body, _err := client.ModifyPlaybookInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes the playbook. After the playbook is published, the playbook runs based on the new logic.
//
// @param request - PublishPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishPlaybookResponse
func (client *Client) PublishPlaybookWithOptions(request *PublishPlaybookRequest, runtime *util.RuntimeOptions) (_result *PublishPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PublishPlaybookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PublishPlaybookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Publishes the playbook. After the playbook is published, the playbook runs based on the new logic.
//
// @param request - PublishPlaybookRequest
//
// @return PublishPlaybookResponse
func (client *Client) PublishPlaybook(request *PublishPlaybookRequest) (_result *PublishPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishPlaybookResponse{}
	_body, _err := client.PublishPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all playbooks at a time.
//
// @param request - QueryTreeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTreeDataResponse
func (client *Client) QueryTreeDataWithOptions(request *QueryTreeDataRequest, runtime *util.RuntimeOptions) (_result *QueryTreeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTreeData"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryTreeDataResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryTreeDataResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries all playbooks at a time.
//
// @param request - QueryTreeDataRequest
//
// @return QueryTreeDataResponse
func (client *Client) QueryTreeData(request *QueryTreeDataRequest) (_result *QueryTreeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTreeDataResponse{}
	_body, _err := client.QueryTreeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name of a node in a playbook. You can call this operation during playbook orchestration. After the name of the node is changed, the reference path of the node also changes.
//
// @param request - RenamePlaybookNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenamePlaybookNodeResponse
func (client *Client) RenamePlaybookNodeWithOptions(request *RenamePlaybookNodeRequest, runtime *util.RuntimeOptions) (_result *RenamePlaybookNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NewNodeName)) {
		query["NewNodeName"] = request.NewNodeName
	}

	if !tea.BoolValue(util.IsUnset(request.OldNodeName)) {
		query["OldNodeName"] = request.OldNodeName
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		query["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenamePlaybookNode"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RenamePlaybookNodeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RenamePlaybookNodeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Changes the name of a node in a playbook. You can call this operation during playbook orchestration. After the name of the node is changed, the reference path of the node also changes.
//
// @param request - RenamePlaybookNodeRequest
//
// @return RenamePlaybookNodeResponse
func (client *Client) RenamePlaybookNode(request *RenamePlaybookNodeRequest) (_result *RenamePlaybookNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenamePlaybookNodeResponse{}
	_body, _err := client.RenamePlaybookNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rolls back a playbook to a specific version. You can determine whether to publish the new playbook version during the rollback.
//
// @param request - RevertPlaybookReleaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertPlaybookReleaseResponse
func (client *Client) RevertPlaybookReleaseWithOptions(request *RevertPlaybookReleaseRequest, runtime *util.RuntimeOptions) (_result *RevertPlaybookReleaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsPublish)) {
		body["IsPublish"] = request.IsPublish
	}

	if !tea.BoolValue(util.IsUnset(request.PlayReleaseId)) {
		body["PlayReleaseId"] = request.PlayReleaseId
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevertPlaybookRelease"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RevertPlaybookReleaseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RevertPlaybookReleaseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Rolls back a playbook to a specific version. You can determine whether to publish the new playbook version during the rollback.
//
// @param request - RevertPlaybookReleaseRequest
//
// @return RevertPlaybookReleaseResponse
func (client *Client) RevertPlaybookRelease(request *RevertPlaybookReleaseRequest) (_result *RevertPlaybookReleaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevertPlaybookReleaseResponse{}
	_body, _err := client.RevertPlaybookReleaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行通知组件-email发送消息
//
// @param request - RunNotifyComponentWithEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunNotifyComponentWithEmailResponse
func (client *Client) RunNotifyComponentWithEmailWithOptions(request *RunNotifyComponentWithEmailRequest, runtime *util.RuntimeOptions) (_result *RunNotifyComponentWithEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionName)) {
		query["ActionName"] = request.ActionName
	}

	if !tea.BoolValue(util.IsUnset(request.AssetId)) {
		query["AssetId"] = request.AssetId
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentName)) {
		query["ComponentName"] = request.ComponentName
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		query["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		query["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Receivers)) {
		query["Receivers"] = request.Receivers
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		query["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		query["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		query["Subject"] = request.Subject
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunNotifyComponentWithEmail"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunNotifyComponentWithEmailResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunNotifyComponentWithEmailResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 执行通知组件-email发送消息
//
// @param request - RunNotifyComponentWithEmailRequest
//
// @return RunNotifyComponentWithEmailResponse
func (client *Client) RunNotifyComponentWithEmail(request *RunNotifyComponentWithEmailRequest) (_result *RunNotifyComponentWithEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunNotifyComponentWithEmailResponse{}
	_body, _err := client.RunNotifyComponentWithEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行通知组件-消息中心发送消息
//
// @param request - RunNotifyComponentWithMessageCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunNotifyComponentWithMessageCenterResponse
func (client *Client) RunNotifyComponentWithMessageCenterWithOptions(request *RunNotifyComponentWithMessageCenterRequest, runtime *util.RuntimeOptions) (_result *RunNotifyComponentWithMessageCenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionName)) {
		query["ActionName"] = request.ActionName
	}

	if !tea.BoolValue(util.IsUnset(request.Aliuid)) {
		query["Aliuid"] = request.Aliuid
	}

	if !tea.BoolValue(util.IsUnset(request.AssetId)) {
		query["AssetId"] = request.AssetId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelTypeList)) {
		query["ChannelTypeList"] = request.ChannelTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentName)) {
		query["ComponentName"] = request.ComponentName
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		query["EventId"] = request.EventId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		query["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		query["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		query["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		query["RoleType"] = request.RoleType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunNotifyComponentWithMessageCenter"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunNotifyComponentWithMessageCenterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunNotifyComponentWithMessageCenterResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 执行通知组件-消息中心发送消息
//
// @param request - RunNotifyComponentWithMessageCenterRequest
//
// @return RunNotifyComponentWithMessageCenterResponse
func (client *Client) RunNotifyComponentWithMessageCenter(request *RunNotifyComponentWithMessageCenterRequest) (_result *RunNotifyComponentWithMessageCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunNotifyComponentWithMessageCenterResponse{}
	_body, _err := client.RunNotifyComponentWithMessageCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行通知组件-webhook发送消息
//
// @param request - RunNotifyComponentWithWebhookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunNotifyComponentWithWebhookResponse
func (client *Client) RunNotifyComponentWithWebhookWithOptions(request *RunNotifyComponentWithWebhookRequest, runtime *util.RuntimeOptions) (_result *RunNotifyComponentWithWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionName)) {
		query["ActionName"] = request.ActionName
	}

	if !tea.BoolValue(util.IsUnset(request.AssetId)) {
		query["AssetId"] = request.AssetId
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentName)) {
		query["ComponentName"] = request.ComponentName
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MsgType)) {
		query["MsgType"] = request.MsgType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		query["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		query["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RoleFor)) {
		query["RoleFor"] = request.RoleFor
	}

	if !tea.BoolValue(util.IsUnset(request.RoleType)) {
		query["RoleType"] = request.RoleType
	}

	if !tea.BoolValue(util.IsUnset(request.Secret)) {
		query["Secret"] = request.Secret
	}

	if !tea.BoolValue(util.IsUnset(request.Webhook)) {
		query["Webhook"] = request.Webhook
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunNotifyComponentWithWebhook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunNotifyComponentWithWebhookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunNotifyComponentWithWebhookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 执行通知组件-webhook发送消息
//
// @param request - RunNotifyComponentWithWebhookRequest
//
// @return RunNotifyComponentWithWebhookResponse
func (client *Client) RunNotifyComponentWithWebhook(request *RunNotifyComponentWithWebhookRequest) (_result *RunNotifyComponentWithWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunNotifyComponentWithWebhookResponse{}
	_body, _err := client.RunNotifyComponentWithWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits and runs a Python3 script. You can call this operation only for data processing.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of Security Orchestration Automation Response (SOAR). For more information, see [Pricing](https://www.alibabacloud.com/en/pricing-calculator?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.4c41281fWhbdPa#/commodity/vm_intl).
//
// @param request - RunPython3ScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunPython3ScriptResponse
func (client *Client) RunPython3ScriptWithOptions(request *RunPython3ScriptRequest, runtime *util.RuntimeOptions) (_result *RunPython3ScriptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		body["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.PythonScript)) {
		body["PythonScript"] = request.PythonScript
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunPython3Script"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunPython3ScriptResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunPython3ScriptResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Submits and runs a Python3 script. You can call this operation only for data processing.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of Security Orchestration Automation Response (SOAR). For more information, see [Pricing](https://www.alibabacloud.com/en/pricing-calculator?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.4c41281fWhbdPa#/commodity/vm_intl).
//
// @param request - RunPython3ScriptRequest
//
// @return RunPython3ScriptResponse
func (client *Client) RunPython3Script(request *RunPython3ScriptRequest) (_result *RunPython3ScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunPython3ScriptResponse{}
	_body, _err := client.RunPython3ScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Triggers an enabled custom playbook or a predefined playbook.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and pricing of Security Orchestration Automation Response (SOAR). For more information, see [Pricing](https://www.alibabacloud.com/en/pricing-calculator?_p_lc=1&spm=a2796.7960336.3034855210.1.7adab91arMeIx2#/commodity/vm_intl).
//
// @param request - TriggerPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerPlaybookResponse
func (client *Client) TriggerPlaybookWithOptions(request *TriggerPlaybookRequest, runtime *util.RuntimeOptions) (_result *TriggerPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InputParam)) {
		body["InputParam"] = request.InputParam
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TriggerPlaybookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TriggerPlaybookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Triggers an enabled custom playbook or a predefined playbook.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and pricing of Security Orchestration Automation Response (SOAR). For more information, see [Pricing](https://www.alibabacloud.com/en/pricing-calculator?_p_lc=1&spm=a2796.7960336.3034855210.1.7adab91arMeIx2#/commodity/vm_intl).
//
// @param request - TriggerPlaybookRequest
//
// @return TriggerPlaybookResponse
func (client *Client) TriggerPlaybook(request *TriggerPlaybookRequest) (_result *TriggerPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerPlaybookResponse{}
	_body, _err := client.TriggerPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs an action on a handling task that is generated by the handling center when an event is handled by using Security Orchestration Automation Response (SOAR). For example, you can call this operation to cancel blocking or isolation, or retry blocking.
//
// @param request - TriggerProcessTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerProcessTaskResponse
func (client *Client) TriggerProcessTaskWithOptions(request *TriggerProcessTaskRequest, runtime *util.RuntimeOptions) (_result *TriggerProcessTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		query["ActionType"] = request.ActionType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerProcessTask"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TriggerProcessTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TriggerProcessTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Performs an action on a handling task that is generated by the handling center when an event is handled by using Security Orchestration Automation Response (SOAR). For example, you can call this operation to cancel blocking or isolation, or retry blocking.
//
// @param request - TriggerProcessTaskRequest
//
// @return TriggerProcessTaskResponse
func (client *Client) TriggerProcessTask(request *TriggerProcessTaskRequest) (_result *TriggerProcessTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerProcessTaskResponse{}
	_body, _err := client.TriggerProcessTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Triggers a playbook or a command.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and pricing of Security Orchestration Automation Response (SOAR). For more information, see [Pricing](https://www.alibabacloud.com/en/pricing-calculator?_p_lc=1&spm=a2796.7960336.3034855210.1.7adab91arMeIx2#/commodity/vm_intl).
//
// @param request - TriggerSophonPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerSophonPlaybookResponse
func (client *Client) TriggerSophonPlaybookWithOptions(request *TriggerSophonPlaybookRequest, runtime *util.RuntimeOptions) (_result *TriggerSophonPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandName)) {
		query["CommandName"] = request.CommandName
	}

	if !tea.BoolValue(util.IsUnset(request.InputParams)) {
		query["InputParams"] = request.InputParams
	}

	if !tea.BoolValue(util.IsUnset(request.SophonTaskId)) {
		query["SophonTaskId"] = request.SophonTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerType)) {
		query["TriggerType"] = request.TriggerType
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerSophonPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TriggerSophonPlaybookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TriggerSophonPlaybookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Triggers a playbook or a command.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and pricing of Security Orchestration Automation Response (SOAR). For more information, see [Pricing](https://www.alibabacloud.com/en/pricing-calculator?_p_lc=1&spm=a2796.7960336.3034855210.1.7adab91arMeIx2#/commodity/vm_intl).
//
// @param request - TriggerSophonPlaybookRequest
//
// @return TriggerSophonPlaybookResponse
func (client *Client) TriggerSophonPlaybook(request *TriggerSophonPlaybookRequest) (_result *TriggerSophonPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerSophonPlaybookResponse{}
	_body, _err := client.TriggerSophonPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the configuration of the playbook is correct and whether the logic of the orchestration is reasonable.
//
// @param request - VerifyPlaybookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyPlaybookResponse
func (client *Client) VerifyPlaybookWithOptions(request *VerifyPlaybookRequest, runtime *util.RuntimeOptions) (_result *VerifyPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.TaskFlow)) {
		body["TaskFlow"] = request.TaskFlow
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyPlaybook"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &VerifyPlaybookResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &VerifyPlaybookResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Checks whether the configuration of the playbook is correct and whether the logic of the orchestration is reasonable.
//
// @param request - VerifyPlaybookRequest
//
// @return VerifyPlaybookResponse
func (client *Client) VerifyPlaybook(request *VerifyPlaybookRequest) (_result *VerifyPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyPlaybookResponse{}
	_body, _err := client.VerifyPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the syntax of a Python code snippet is correct.
//
// @param request - VerifyPythonFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyPythonFileResponse
func (client *Client) VerifyPythonFileWithOptions(request *VerifyPythonFileRequest, runtime *util.RuntimeOptions) (_result *VerifyPythonFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyPythonFile"),
		Version:     tea.String("2022-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &VerifyPythonFileResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &VerifyPythonFileResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Checks whether the syntax of a Python code snippet is correct.
//
// @param request - VerifyPythonFileRequest
//
// @return VerifyPythonFileResponse
func (client *Client) VerifyPythonFile(request *VerifyPythonFileRequest) (_result *VerifyPythonFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyPythonFileResponse{}
	_body, _err := client.VerifyPythonFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
