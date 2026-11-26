// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersonalSourcesByObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListPersonalSourcesByObjectRequest
	GetAppToken() *string
	SetGraphName(v string) *ListPersonalSourcesByObjectRequest
	GetGraphName() *string
	SetObjectId(v string) *ListPersonalSourcesByObjectRequest
	GetObjectId() *string
	SetObjectType(v string) *ListPersonalSourcesByObjectRequest
	GetObjectType() *string
	SetOpsUserCode(v string) *ListPersonalSourcesByObjectRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListPersonalSourcesByObjectRequest
	GetOpsUserId() *string
	SetPage(v int64) *ListPersonalSourcesByObjectRequest
	GetPage() *int64
	SetPageSize(v int64) *ListPersonalSourcesByObjectRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListPersonalSourcesByObjectRequest
	GetTenantId() *string
}

type ListPersonalSourcesByObjectRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// 绑定对象所属图谱名称；与 objectType、objectId 共同精确定位对象
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// 绑定对象记录 ID；后端按对象主键类型规范化后查询
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 绑定对象类型；先检查当前用户对目标对象记录的读权限
	//
	// This parameter is required.
	//
	// example:
	//
	// contract
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// 页码，从 1 开始，默认 1；越界页返回空列表及授权后的总数
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数，范围 1-100，默认 20
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListPersonalSourcesByObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalSourcesByObjectRequest) GoString() string {
	return s.String()
}

func (s *ListPersonalSourcesByObjectRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListPersonalSourcesByObjectRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *ListPersonalSourcesByObjectRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListPersonalSourcesByObjectRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListPersonalSourcesByObjectRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListPersonalSourcesByObjectRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListPersonalSourcesByObjectRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListPersonalSourcesByObjectRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPersonalSourcesByObjectRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListPersonalSourcesByObjectRequest) SetAppToken(v string) *ListPersonalSourcesByObjectRequest {
	s.AppToken = &v
	return s
}

func (s *ListPersonalSourcesByObjectRequest) SetGraphName(v string) *ListPersonalSourcesByObjectRequest {
	s.GraphName = &v
	return s
}

func (s *ListPersonalSourcesByObjectRequest) SetObjectId(v string) *ListPersonalSourcesByObjectRequest {
	s.ObjectId = &v
	return s
}

func (s *ListPersonalSourcesByObjectRequest) SetObjectType(v string) *ListPersonalSourcesByObjectRequest {
	s.ObjectType = &v
	return s
}

func (s *ListPersonalSourcesByObjectRequest) SetOpsUserCode(v string) *ListPersonalSourcesByObjectRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListPersonalSourcesByObjectRequest) SetOpsUserId(v string) *ListPersonalSourcesByObjectRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListPersonalSourcesByObjectRequest) SetPage(v int64) *ListPersonalSourcesByObjectRequest {
	s.Page = &v
	return s
}

func (s *ListPersonalSourcesByObjectRequest) SetPageSize(v int64) *ListPersonalSourcesByObjectRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonalSourcesByObjectRequest) SetTenantId(v string) *ListPersonalSourcesByObjectRequest {
	s.TenantId = &v
	return s
}

func (s *ListPersonalSourcesByObjectRequest) Validate() error {
	return dara.Validate(s)
}
