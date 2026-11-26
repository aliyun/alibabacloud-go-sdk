// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceObjectBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ReplaceObjectBindingsRequest
	GetAppToken() *string
	SetObjectBindings(v []*ReplaceObjectBindingsRequestObjectBindings) *ReplaceObjectBindingsRequest
	GetObjectBindings() []*ReplaceObjectBindingsRequestObjectBindings
	SetOpsUserCode(v string) *ReplaceObjectBindingsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ReplaceObjectBindingsRequest
	GetOpsUserId() *string
	SetSourceId(v string) *ReplaceObjectBindingsRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplaceObjectBindingsRequest
	GetTenantId() *string
}

type ReplaceObjectBindingsRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The new list of object bindings for full replacement. Passing an empty list clears all bindings.
	//
	// This parameter is required.
	ObjectBindings []*ReplaceObjectBindingsRequestObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
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
	// The ID of the personal FILE data source to be replaced, which is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter and is explicitly passed by winnexo-cli through the --tenant-id option.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReplaceObjectBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceObjectBindingsRequest) GoString() string {
	return s.String()
}

func (s *ReplaceObjectBindingsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ReplaceObjectBindingsRequest) GetObjectBindings() []*ReplaceObjectBindingsRequestObjectBindings {
	return s.ObjectBindings
}

func (s *ReplaceObjectBindingsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ReplaceObjectBindingsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ReplaceObjectBindingsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceObjectBindingsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceObjectBindingsRequest) SetAppToken(v string) *ReplaceObjectBindingsRequest {
	s.AppToken = &v
	return s
}

func (s *ReplaceObjectBindingsRequest) SetObjectBindings(v []*ReplaceObjectBindingsRequestObjectBindings) *ReplaceObjectBindingsRequest {
	s.ObjectBindings = v
	return s
}

func (s *ReplaceObjectBindingsRequest) SetOpsUserCode(v string) *ReplaceObjectBindingsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ReplaceObjectBindingsRequest) SetOpsUserId(v string) *ReplaceObjectBindingsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ReplaceObjectBindingsRequest) SetSourceId(v string) *ReplaceObjectBindingsRequest {
	s.SourceId = &v
	return s
}

func (s *ReplaceObjectBindingsRequest) SetTenantId(v string) *ReplaceObjectBindingsRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceObjectBindingsRequest) Validate() error {
	if s.ObjectBindings != nil {
		for _, item := range s.ObjectBindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReplaceObjectBindingsRequestObjectBindings struct {
	// The name of the semantic graph to which the bound object belongs. The object_id must be unique within this graph. This parameter is required.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The ID of the bound object.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object type to attach, such as customer or project.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s ReplaceObjectBindingsRequestObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s ReplaceObjectBindingsRequestObjectBindings) GoString() string {
	return s.String()
}

func (s *ReplaceObjectBindingsRequestObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *ReplaceObjectBindingsRequestObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *ReplaceObjectBindingsRequestObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *ReplaceObjectBindingsRequestObjectBindings) SetGraphName(v string) *ReplaceObjectBindingsRequestObjectBindings {
	s.GraphName = &v
	return s
}

func (s *ReplaceObjectBindingsRequestObjectBindings) SetObjectId(v string) *ReplaceObjectBindingsRequestObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *ReplaceObjectBindingsRequestObjectBindings) SetObjectType(v string) *ReplaceObjectBindingsRequestObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *ReplaceObjectBindingsRequestObjectBindings) Validate() error {
	return dara.Validate(s)
}
