// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemoveOperatingObjectFavoritesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetAppToken() *string
	SetGraphName(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetGraphName() *string
	SetObjectIdsShrink(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetObjectIdsShrink() *string
	SetObjectType(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetObjectType() *string
	SetOperatingObjectName(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetTenantId() *string
}

type BatchRemoveOperatingObjectFavoritesShrinkRequest struct {
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
	ObjectIdsShrink *string `json:"objectIds,omitempty" xml:"objectIds,omitempty"`
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

func (s BatchRemoveOperatingObjectFavoritesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRemoveOperatingObjectFavoritesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetObjectIdsShrink() *string {
	return s.ObjectIdsShrink
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetAppToken(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetGraphName(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.GraphName = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetObjectIdsShrink(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.ObjectIdsShrink = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetObjectType(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.ObjectType = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetOperatingObjectName(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetOpsUserCode(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetOpsUserId(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetTenantId(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
