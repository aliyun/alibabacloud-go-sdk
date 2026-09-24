// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageQueryAgentListNewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *PageQueryAgentListNewResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *PageQueryAgentListNewResponseBody
	GetCode() *string
	SetData(v *PageQueryAgentListNewResponseBodyData) *PageQueryAgentListNewResponseBody
	GetData() *PageQueryAgentListNewResponseBodyData
	SetMessage(v string) *PageQueryAgentListNewResponseBody
	GetMessage() *string
	SetRequestId(v string) *PageQueryAgentListNewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PageQueryAgentListNewResponseBody
	GetSuccess() *bool
}

type PageQueryAgentListNewResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// Access denied due to insufficient permissions
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *PageQueryAgentListNewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code description.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful.
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PageQueryAgentListNewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PageQueryAgentListNewResponseBody) GoString() string {
	return s.String()
}

func (s *PageQueryAgentListNewResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *PageQueryAgentListNewResponseBody) GetCode() *string {
	return s.Code
}

func (s *PageQueryAgentListNewResponseBody) GetData() *PageQueryAgentListNewResponseBodyData {
	return s.Data
}

func (s *PageQueryAgentListNewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PageQueryAgentListNewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PageQueryAgentListNewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PageQueryAgentListNewResponseBody) SetAccessDeniedDetail(v string) *PageQueryAgentListNewResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *PageQueryAgentListNewResponseBody) SetCode(v string) *PageQueryAgentListNewResponseBody {
	s.Code = &v
	return s
}

func (s *PageQueryAgentListNewResponseBody) SetData(v *PageQueryAgentListNewResponseBodyData) *PageQueryAgentListNewResponseBody {
	s.Data = v
	return s
}

func (s *PageQueryAgentListNewResponseBody) SetMessage(v string) *PageQueryAgentListNewResponseBody {
	s.Message = &v
	return s
}

func (s *PageQueryAgentListNewResponseBody) SetRequestId(v string) *PageQueryAgentListNewResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageQueryAgentListNewResponseBody) SetSuccess(v bool) *PageQueryAgentListNewResponseBody {
	s.Success = &v
	return s
}

func (s *PageQueryAgentListNewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PageQueryAgentListNewResponseBodyData struct {
	// The data list.
	List []*PageQueryAgentListNewResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 28
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The page size.
	//
	// example:
	//
	// 41
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 6
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PageQueryAgentListNewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PageQueryAgentListNewResponseBodyData) GoString() string {
	return s.String()
}

func (s *PageQueryAgentListNewResponseBodyData) GetList() []*PageQueryAgentListNewResponseBodyDataList {
	return s.List
}

func (s *PageQueryAgentListNewResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *PageQueryAgentListNewResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *PageQueryAgentListNewResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *PageQueryAgentListNewResponseBodyData) SetList(v []*PageQueryAgentListNewResponseBodyDataList) *PageQueryAgentListNewResponseBodyData {
	s.List = v
	return s
}

func (s *PageQueryAgentListNewResponseBodyData) SetPageNo(v int64) *PageQueryAgentListNewResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyData) SetPageSize(v int64) *PageQueryAgentListNewResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyData) SetTotal(v int64) *PageQueryAgentListNewResponseBodyData {
	s.Total = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PageQueryAgentListNewResponseBodyDataList struct {
	// The agent ID.
	//
	// example:
	//
	// 51
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The agent building mode. Valid values:
	//
	// - 0: prompt mode (PROMPT).
	//
	// - 1: dialog flow mode (CONVERSATION).
	//
	// example:
	//
	// 0
	AgentMode *int64 `json:"AgentMode,omitempty" xml:"AgentMode,omitempty"`
	// The agent name.
	//
	// example:
	//
	// Intelligent Customer Service Assistant
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The application code.
	//
	// example:
	//
	// aicc_demo_app
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// The creation time, in the format of YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2024-01-20 12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the branch being deployed.
	//
	// example:
	//
	// 24
	DeployBranchId *int64 `json:"DeployBranchId,omitempty" xml:"DeployBranchId,omitempty"`
	// The name of the active branch.
	//
	// example:
	//
	// master
	DeployBranchName *string `json:"DeployBranchName,omitempty" xml:"DeployBranchName,omitempty"`
	// The agent description.
	//
	// example:
	//
	// Intelligent customer service assistant that provides automated customer service support
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the active version.
	//
	// example:
	//
	// 71
	EffectiveVersionId *int64 `json:"EffectiveVersionId,omitempty" xml:"EffectiveVersionId,omitempty"`
	// The name of the active version.
	//
	// example:
	//
	// v1.0.0
	EffectiveVersionName *string `json:"EffectiveVersionName,omitempty" xml:"EffectiveVersionName,omitempty"`
	// Indicates whether the agent is available for outbound calls. A value of True indicates that the current deployment branch of the agent has a published version and is available for outbound calls.
	//
	// example:
	//
	// true
	IsAvailable *bool `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	// The latest version publish time, in the format of YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2024-01-20 12:00:00
	LatestPublishTime *string `json:"LatestPublishTime,omitempty" xml:"LatestPublishTime,omitempty"`
	// The last modification time, in the format of YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2024-01-15 10:30:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The scenario.
	//
	// example:
	//
	// Personal lead conversion
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The service direction.
	//
	// example:
	//
	// Sample value
	ServiceDirection *string `json:"ServiceDirection,omitempty" xml:"ServiceDirection,omitempty"`
	// The source template ID.
	//
	// example:
	//
	// 62
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The source template name.
	//
	// example:
	//
	// Sample value
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s PageQueryAgentListNewResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s PageQueryAgentListNewResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetAgentId() *int64 {
	return s.AgentId
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetAgentMode() *int64 {
	return s.AgentMode
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetDeployBranchId() *int64 {
	return s.DeployBranchId
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetDeployBranchName() *string {
	return s.DeployBranchName
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetEffectiveVersionId() *int64 {
	return s.EffectiveVersionId
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetEffectiveVersionName() *string {
	return s.EffectiveVersionName
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetIsAvailable() *bool {
	return s.IsAvailable
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetLatestPublishTime() *string {
	return s.LatestPublishTime
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetScene() *string {
	return s.Scene
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetServiceDirection() *string {
	return s.ServiceDirection
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *PageQueryAgentListNewResponseBodyDataList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetAgentId(v int64) *PageQueryAgentListNewResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetAgentMode(v int64) *PageQueryAgentListNewResponseBodyDataList {
	s.AgentMode = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetAgentName(v string) *PageQueryAgentListNewResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetApplicationCode(v string) *PageQueryAgentListNewResponseBodyDataList {
	s.ApplicationCode = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetCreateTime(v string) *PageQueryAgentListNewResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetDeployBranchId(v int64) *PageQueryAgentListNewResponseBodyDataList {
	s.DeployBranchId = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetDeployBranchName(v string) *PageQueryAgentListNewResponseBodyDataList {
	s.DeployBranchName = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetDescription(v string) *PageQueryAgentListNewResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetEffectiveVersionId(v int64) *PageQueryAgentListNewResponseBodyDataList {
	s.EffectiveVersionId = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetEffectiveVersionName(v string) *PageQueryAgentListNewResponseBodyDataList {
	s.EffectiveVersionName = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetIsAvailable(v bool) *PageQueryAgentListNewResponseBodyDataList {
	s.IsAvailable = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetLatestPublishTime(v string) *PageQueryAgentListNewResponseBodyDataList {
	s.LatestPublishTime = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetModifyTime(v string) *PageQueryAgentListNewResponseBodyDataList {
	s.ModifyTime = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetScene(v string) *PageQueryAgentListNewResponseBodyDataList {
	s.Scene = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetServiceDirection(v string) *PageQueryAgentListNewResponseBodyDataList {
	s.ServiceDirection = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetTemplateId(v int64) *PageQueryAgentListNewResponseBodyDataList {
	s.TemplateId = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) SetTemplateName(v string) *PageQueryAgentListNewResponseBodyDataList {
	s.TemplateName = &v
	return s
}

func (s *PageQueryAgentListNewResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
