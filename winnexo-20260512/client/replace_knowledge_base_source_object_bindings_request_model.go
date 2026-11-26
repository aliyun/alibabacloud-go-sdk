// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceKnowledgeBaseSourceObjectBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequest
	GetAppToken() *string
	SetObjectBindings(v []*ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings) *ReplaceKnowledgeBaseSourceObjectBindingsRequest
	GetObjectBindings() []*ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings
	SetOpsUserCode(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequest
	GetOpsUserId() *string
	SetSourceId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequest
	GetTenantId() *string
}

type ReplaceKnowledgeBaseSourceObjectBindingsRequest struct {
	// token
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The object bindings.
	//
	// This parameter is required.
	ObjectBindings []*ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
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
	// The ID of the data source to be re-parsed. The ID is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. For winnexo-cli, pass this parameter explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReplaceKnowledgeBaseSourceObjectBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseSourceObjectBindingsRequest) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) GetObjectBindings() []*ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings {
	return s.ObjectBindings
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) SetAppToken(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequest {
	s.AppToken = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) SetObjectBindings(v []*ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings) *ReplaceKnowledgeBaseSourceObjectBindingsRequest {
	s.ObjectBindings = v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) SetOpsUserCode(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) SetOpsUserId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) SetSourceId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequest {
	s.SourceId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) SetTenantId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequest) Validate() error {
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

type ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings struct {
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

func (s ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings) SetGraphName(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings {
	s.GraphName = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings) SetObjectId(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings) SetObjectType(v string) *ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceObjectBindingsRequestObjectBindings) Validate() error {
	return dara.Validate(s)
}
