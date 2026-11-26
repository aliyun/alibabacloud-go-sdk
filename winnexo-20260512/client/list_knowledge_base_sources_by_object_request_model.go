// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseSourcesByObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListKnowledgeBaseSourcesByObjectRequest
	GetAppToken() *string
	SetGraphName(v string) *ListKnowledgeBaseSourcesByObjectRequest
	GetGraphName() *string
	SetObjectId(v string) *ListKnowledgeBaseSourcesByObjectRequest
	GetObjectId() *string
	SetObjectType(v string) *ListKnowledgeBaseSourcesByObjectRequest
	GetObjectType() *string
	SetOpsUserCode(v string) *ListKnowledgeBaseSourcesByObjectRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListKnowledgeBaseSourcesByObjectRequest
	GetOpsUserId() *string
	SetPage(v int64) *ListKnowledgeBaseSourcesByObjectRequest
	GetPage() *int64
	SetPageSize(v int64) *ListKnowledgeBaseSourcesByObjectRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListKnowledgeBaseSourcesByObjectRequest
	GetTenantId() *string
}

type ListKnowledgeBaseSourcesByObjectRequest struct {
	// token
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The graph name. You can call listGraphs to obtain it.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The ID of the associated variable.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2676
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// This parameter is required.
	//
	// example:
	//
	// contract
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The code of the operator.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, it is explicitly passed by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListKnowledgeBaseSourcesByObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseSourcesByObjectRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) SetAppToken(v string) *ListKnowledgeBaseSourcesByObjectRequest {
	s.AppToken = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) SetGraphName(v string) *ListKnowledgeBaseSourcesByObjectRequest {
	s.GraphName = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) SetObjectId(v string) *ListKnowledgeBaseSourcesByObjectRequest {
	s.ObjectId = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) SetObjectType(v string) *ListKnowledgeBaseSourcesByObjectRequest {
	s.ObjectType = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) SetOpsUserCode(v string) *ListKnowledgeBaseSourcesByObjectRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) SetOpsUserId(v string) *ListKnowledgeBaseSourcesByObjectRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) SetPage(v int64) *ListKnowledgeBaseSourcesByObjectRequest {
	s.Page = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) SetPageSize(v int64) *ListKnowledgeBaseSourcesByObjectRequest {
	s.PageSize = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) SetTenantId(v string) *ListKnowledgeBaseSourcesByObjectRequest {
	s.TenantId = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectRequest) Validate() error {
	return dara.Validate(s)
}
