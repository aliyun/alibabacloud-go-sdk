// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplacePersonalSourceObjectBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ReplacePersonalSourceObjectBindingsRequest
	GetAppToken() *string
	SetObjectBindings(v []*ReplacePersonalSourceObjectBindingsRequestObjectBindings) *ReplacePersonalSourceObjectBindingsRequest
	GetObjectBindings() []*ReplacePersonalSourceObjectBindingsRequestObjectBindings
	SetOpsUserCode(v string) *ReplacePersonalSourceObjectBindingsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ReplacePersonalSourceObjectBindingsRequest
	GetOpsUserId() *string
	SetSourceId(v string) *ReplacePersonalSourceObjectBindingsRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplacePersonalSourceObjectBindingsRequest
	GetTenantId() *string
}

type ReplacePersonalSourceObjectBindingsRequest struct {
	// token
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The object bindings.
	//
	// This parameter is required.
	ObjectBindings []*ReplacePersonalSourceObjectBindingsRequestObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
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
	// The ID of the data source. The ID is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReplacePersonalSourceObjectBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplacePersonalSourceObjectBindingsRequest) GoString() string {
	return s.String()
}

func (s *ReplacePersonalSourceObjectBindingsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ReplacePersonalSourceObjectBindingsRequest) GetObjectBindings() []*ReplacePersonalSourceObjectBindingsRequestObjectBindings {
	return s.ObjectBindings
}

func (s *ReplacePersonalSourceObjectBindingsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ReplacePersonalSourceObjectBindingsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ReplacePersonalSourceObjectBindingsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplacePersonalSourceObjectBindingsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplacePersonalSourceObjectBindingsRequest) SetAppToken(v string) *ReplacePersonalSourceObjectBindingsRequest {
	s.AppToken = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsRequest) SetObjectBindings(v []*ReplacePersonalSourceObjectBindingsRequestObjectBindings) *ReplacePersonalSourceObjectBindingsRequest {
	s.ObjectBindings = v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsRequest) SetOpsUserCode(v string) *ReplacePersonalSourceObjectBindingsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsRequest) SetOpsUserId(v string) *ReplacePersonalSourceObjectBindingsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsRequest) SetSourceId(v string) *ReplacePersonalSourceObjectBindingsRequest {
	s.SourceId = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsRequest) SetTenantId(v string) *ReplacePersonalSourceObjectBindingsRequest {
	s.TenantId = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsRequest) Validate() error {
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

type ReplacePersonalSourceObjectBindingsRequestObjectBindings struct {
	// The name of the semantic graph to which the bound object belongs. The object ID is unique within this graph. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The ID of the bound object.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object type to attach, such as customer or project.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s ReplacePersonalSourceObjectBindingsRequestObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s ReplacePersonalSourceObjectBindingsRequestObjectBindings) GoString() string {
	return s.String()
}

func (s *ReplacePersonalSourceObjectBindingsRequestObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *ReplacePersonalSourceObjectBindingsRequestObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *ReplacePersonalSourceObjectBindingsRequestObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *ReplacePersonalSourceObjectBindingsRequestObjectBindings) SetGraphName(v string) *ReplacePersonalSourceObjectBindingsRequestObjectBindings {
	s.GraphName = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsRequestObjectBindings) SetObjectId(v string) *ReplacePersonalSourceObjectBindingsRequestObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsRequestObjectBindings) SetObjectType(v string) *ReplacePersonalSourceObjectBindingsRequestObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *ReplacePersonalSourceObjectBindingsRequestObjectBindings) Validate() error {
	return dara.Validate(s)
}
