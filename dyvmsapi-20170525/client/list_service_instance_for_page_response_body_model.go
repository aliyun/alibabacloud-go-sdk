// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceForPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListServiceInstanceForPageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListServiceInstanceForPageResponseBody
	GetCode() *string
	SetMessage(v string) *ListServiceInstanceForPageResponseBody
	GetMessage() *string
	SetModel(v *ListServiceInstanceForPageResponseBodyModel) *ListServiceInstanceForPageResponseBody
	GetModel() *ListServiceInstanceForPageResponseBodyModel
	SetRequestId(v string) *ListServiceInstanceForPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListServiceInstanceForPageResponseBody
	GetSuccess() *bool
}

type ListServiceInstanceForPageResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 状态码，OK表示响应成功
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 响应消息
	//
	// example:
	//
	// 每页最大100条！
	Message *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *ListServiceInstanceForPageResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// 请求唯一id
	//
	// example:
	//
	// ED18A5AE-9BBC-5851-1111-8E5B8C23CEDE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListServiceInstanceForPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceForPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceForPageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListServiceInstanceForPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListServiceInstanceForPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListServiceInstanceForPageResponseBody) GetModel() *ListServiceInstanceForPageResponseBodyModel {
	return s.Model
}

func (s *ListServiceInstanceForPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceInstanceForPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListServiceInstanceForPageResponseBody) SetAccessDeniedDetail(v string) *ListServiceInstanceForPageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBody) SetCode(v string) *ListServiceInstanceForPageResponseBody {
	s.Code = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBody) SetMessage(v string) *ListServiceInstanceForPageResponseBody {
	s.Message = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBody) SetModel(v *ListServiceInstanceForPageResponseBodyModel) *ListServiceInstanceForPageResponseBody {
	s.Model = v
	return s
}

func (s *ListServiceInstanceForPageResponseBody) SetRequestId(v string) *ListServiceInstanceForPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBody) SetSuccess(v bool) *ListServiceInstanceForPageResponseBody {
	s.Success = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServiceInstanceForPageResponseBodyModel struct {
	// 当前页码
	//
	// example:
	//
	// 47
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 每页数量
	//
	// example:
	//
	// 59
	PageSize *int64                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*ListServiceInstanceForPageResponseBodyModelRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// 总数
	//
	// example:
	//
	// 36
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 总页数
	//
	// example:
	//
	// 76
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListServiceInstanceForPageResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceForPageResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceForPageResponseBodyModel) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListServiceInstanceForPageResponseBodyModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListServiceInstanceForPageResponseBodyModel) GetRecords() []*ListServiceInstanceForPageResponseBodyModelRecords {
	return s.Records
}

func (s *ListServiceInstanceForPageResponseBodyModel) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListServiceInstanceForPageResponseBodyModel) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *ListServiceInstanceForPageResponseBodyModel) SetPageNo(v int64) *ListServiceInstanceForPageResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModel) SetPageSize(v int64) *ListServiceInstanceForPageResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModel) SetRecords(v []*ListServiceInstanceForPageResponseBodyModelRecords) *ListServiceInstanceForPageResponseBodyModel {
	s.Records = v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModel) SetTotalCount(v int64) *ListServiceInstanceForPageResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModel) SetTotalPage(v int64) *ListServiceInstanceForPageResponseBodyModel {
	s.TotalPage = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModel) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceInstanceForPageResponseBodyModelRecords struct {
	// aliUid
	//
	// example:
	//
	// 11******123
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 服务实例号
	//
	// example:
	//
	// 0571****1234
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 创建时间（时间戳）
	//
	// example:
	//
	// 1759109856000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 实例状态（init：初始化；usable：可用；unusable：不可用；destory：注销）
	//
	// example:
	//
	// init
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// 服务实例名称
	//
	// example:
	//
	// 测试数据
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 关联可用数量
	//
	// example:
	//
	// 30
	RelationAvailableCount *int64 `json:"RelationAvailableCount,omitempty" xml:"RelationAvailableCount,omitempty"`
	// 关联总数量
	//
	// example:
	//
	// 33
	RelationTotalCount *int64 `json:"RelationTotalCount,omitempty" xml:"RelationTotalCount,omitempty"`
	// 场景ID
	//
	// example:
	//
	// 29
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 场景名称
	//
	// example:
	//
	// 系统通知
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// 绑定状态：绑定；未绑定
	//
	// example:
	//
	// 绑定
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用途ID
	//
	// example:
	//
	// 96
	UsageId *int64 `json:"UsageId,omitempty" xml:"UsageId,omitempty"`
	// 用途名称
	//
	// example:
	//
	// 语音通知
	UsageName *string `json:"UsageName,omitempty" xml:"UsageName,omitempty"`
}

func (s ListServiceInstanceForPageResponseBodyModelRecords) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceForPageResponseBodyModelRecords) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) GetCode() *string {
	return s.Code
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) GetName() *string {
	return s.Name
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) GetRelationAvailableCount() *int64 {
	return s.RelationAvailableCount
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) GetRelationTotalCount() *int64 {
	return s.RelationTotalCount
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) GetSceneId() *int64 {
	return s.SceneId
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) GetSceneName() *string {
	return s.SceneName
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) GetStatus() *string {
	return s.Status
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) GetUsageId() *int64 {
	return s.UsageId
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) GetUsageName() *string {
	return s.UsageName
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) SetAliUid(v int64) *ListServiceInstanceForPageResponseBodyModelRecords {
	s.AliUid = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) SetCode(v string) *ListServiceInstanceForPageResponseBodyModelRecords {
	s.Code = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) SetGmtCreate(v string) *ListServiceInstanceForPageResponseBodyModelRecords {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) SetInstanceStatus(v string) *ListServiceInstanceForPageResponseBodyModelRecords {
	s.InstanceStatus = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) SetName(v string) *ListServiceInstanceForPageResponseBodyModelRecords {
	s.Name = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) SetRelationAvailableCount(v int64) *ListServiceInstanceForPageResponseBodyModelRecords {
	s.RelationAvailableCount = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) SetRelationTotalCount(v int64) *ListServiceInstanceForPageResponseBodyModelRecords {
	s.RelationTotalCount = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) SetSceneId(v int64) *ListServiceInstanceForPageResponseBodyModelRecords {
	s.SceneId = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) SetSceneName(v string) *ListServiceInstanceForPageResponseBodyModelRecords {
	s.SceneName = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) SetStatus(v string) *ListServiceInstanceForPageResponseBodyModelRecords {
	s.Status = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) SetUsageId(v int64) *ListServiceInstanceForPageResponseBodyModelRecords {
	s.UsageId = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) SetUsageName(v string) *ListServiceInstanceForPageResponseBodyModelRecords {
	s.UsageName = &v
	return s
}

func (s *ListServiceInstanceForPageResponseBodyModelRecords) Validate() error {
	return dara.Validate(s)
}
