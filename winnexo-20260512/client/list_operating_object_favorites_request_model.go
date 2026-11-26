// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperatingObjectFavoritesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListOperatingObjectFavoritesRequest
	GetAppToken() *string
	SetGraphName(v string) *ListOperatingObjectFavoritesRequest
	GetGraphName() *string
	SetNextToken(v string) *ListOperatingObjectFavoritesRequest
	GetNextToken() *string
	SetObjectType(v string) *ListOperatingObjectFavoritesRequest
	GetObjectType() *string
	SetOperatingObjectName(v string) *ListOperatingObjectFavoritesRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *ListOperatingObjectFavoritesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListOperatingObjectFavoritesRequest
	GetOpsUserId() *string
	SetPageSize(v int64) *ListOperatingObjectFavoritesRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListOperatingObjectFavoritesRequest
	GetTenantId() *string
}

type ListOperatingObjectFavoritesRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The name of the graph. You can call the listGraphs operation to obtain the graph name.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The pagination cursor.
	//
	// example:
	//
	// eyJ2IjoxLCJpZCI6MTAwMX0.c2lnbmF0dXJlX2V4YW1wbGU
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// This parameter is required.
	//
	// example:
	//
	// contract
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The name of the digital worker, which is the name of the operating object. This parameter is optional.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer_assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListOperatingObjectFavoritesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperatingObjectFavoritesRequest) GoString() string {
	return s.String()
}

func (s *ListOperatingObjectFavoritesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListOperatingObjectFavoritesRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *ListOperatingObjectFavoritesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOperatingObjectFavoritesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListOperatingObjectFavoritesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListOperatingObjectFavoritesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListOperatingObjectFavoritesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListOperatingObjectFavoritesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOperatingObjectFavoritesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListOperatingObjectFavoritesRequest) SetAppToken(v string) *ListOperatingObjectFavoritesRequest {
	s.AppToken = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetGraphName(v string) *ListOperatingObjectFavoritesRequest {
	s.GraphName = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetNextToken(v string) *ListOperatingObjectFavoritesRequest {
	s.NextToken = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetObjectType(v string) *ListOperatingObjectFavoritesRequest {
	s.ObjectType = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetOperatingObjectName(v string) *ListOperatingObjectFavoritesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetOpsUserCode(v string) *ListOperatingObjectFavoritesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetOpsUserId(v string) *ListOperatingObjectFavoritesRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetPageSize(v int64) *ListOperatingObjectFavoritesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) SetTenantId(v string) *ListOperatingObjectFavoritesRequest {
	s.TenantId = &v
	return s
}

func (s *ListOperatingObjectFavoritesRequest) Validate() error {
	return dara.Validate(s)
}
