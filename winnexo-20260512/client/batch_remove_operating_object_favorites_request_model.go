// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemoveOperatingObjectFavoritesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *BatchRemoveOperatingObjectFavoritesRequest
	GetAppToken() *string
	SetGraphName(v string) *BatchRemoveOperatingObjectFavoritesRequest
	GetGraphName() *string
	SetObjectIds(v []*string) *BatchRemoveOperatingObjectFavoritesRequest
	GetObjectIds() []*string
	SetObjectType(v string) *BatchRemoveOperatingObjectFavoritesRequest
	GetObjectType() *string
	SetOperatingObjectName(v string) *BatchRemoveOperatingObjectFavoritesRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *BatchRemoveOperatingObjectFavoritesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *BatchRemoveOperatingObjectFavoritesRequest
	GetOpsUserId() *string
	SetTenantId(v string) *BatchRemoveOperatingObjectFavoritesRequest
	GetTenantId() *string
}

type BatchRemoveOperatingObjectFavoritesRequest struct {
	// The platform application token issued by the console. If provided, the operation is executed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The graph name.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The list of business IDs for the primary objects.
	//
	// This parameter is required.
	//
	// example:
	//
	// contract-001
	ObjectIds []*string `json:"objectIds,omitempty" xml:"objectIds,omitempty" type:"Repeated"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// This parameter is required.
	//
	// example:
	//
	// contract
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The operating object name.
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
	// The tenant ID. This is a common parameter. For winnexo-cli, it is explicitly passed by using the --tenant-id option.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s BatchRemoveOperatingObjectFavoritesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRemoveOperatingObjectFavoritesRequest) GoString() string {
	return s.String()
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetObjectIds() []*string {
	return s.ObjectIds
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetAppToken(v string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.AppToken = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetGraphName(v string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.GraphName = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetObjectIds(v []*string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.ObjectIds = v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetObjectType(v string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.ObjectType = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetOperatingObjectName(v string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetOpsUserCode(v string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetOpsUserId(v string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.OpsUserId = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) SetTenantId(v string) *BatchRemoveOperatingObjectFavoritesRequest {
	s.TenantId = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesRequest) Validate() error {
	return dara.Validate(s)
}
