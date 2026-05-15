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
	// example:
	//
	// Access denied due to insufficient permissions
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *PageQueryAgentListNewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 12345678-1234-1234-1234-123456789012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	List []*PageQueryAgentListNewResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 28
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 41
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// example:
	//
	// 51
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 0
	AgentMode *int64 `json:"AgentMode,omitempty" xml:"AgentMode,omitempty"`
	// example:
	//
	// 示例值示例值
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// aicc_demo_app
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// example:
	//
	// 2024-01-20 12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 24
	DeployBranchId *int64 `json:"DeployBranchId,omitempty" xml:"DeployBranchId,omitempty"`
	// example:
	//
	// master
	DeployBranchName *string `json:"DeployBranchName,omitempty" xml:"DeployBranchName,omitempty"`
	// example:
	//
	// 示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 71
	EffectiveVersionId *int64 `json:"EffectiveVersionId,omitempty" xml:"EffectiveVersionId,omitempty"`
	// example:
	//
	// v1.0.0
	EffectiveVersionName *string `json:"EffectiveVersionName,omitempty" xml:"EffectiveVersionName,omitempty"`
	// example:
	//
	// true
	IsAvailable *bool `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	// example:
	//
	// 2024-01-20 12:00:00
	LatestPublishTime *string `json:"LatestPublishTime,omitempty" xml:"LatestPublishTime,omitempty"`
	// example:
	//
	// 2024-01-15 10:30:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 示例值示例值
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
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

func (s *PageQueryAgentListNewResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
