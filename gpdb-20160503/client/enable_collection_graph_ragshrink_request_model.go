// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCollectionGraphRAGShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCollection(v string) *EnableCollectionGraphRAGShrinkRequest
  GetCollection() *string 
  SetDBInstanceId(v string) *EnableCollectionGraphRAGShrinkRequest
  GetDBInstanceId() *string 
  SetEntityTypesShrink(v string) *EnableCollectionGraphRAGShrinkRequest
  GetEntityTypesShrink() *string 
  SetLLMModel(v string) *EnableCollectionGraphRAGShrinkRequest
  GetLLMModel() *string 
  SetLanguage(v string) *EnableCollectionGraphRAGShrinkRequest
  GetLanguage() *string 
  SetManagerAccount(v string) *EnableCollectionGraphRAGShrinkRequest
  GetManagerAccount() *string 
  SetManagerAccountPassword(v string) *EnableCollectionGraphRAGShrinkRequest
  GetManagerAccountPassword() *string 
  SetNamespace(v string) *EnableCollectionGraphRAGShrinkRequest
  GetNamespace() *string 
  SetNamespacePassword(v string) *EnableCollectionGraphRAGShrinkRequest
  GetNamespacePassword() *string 
  SetOwnerId(v int64) *EnableCollectionGraphRAGShrinkRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableCollectionGraphRAGShrinkRequest
  GetRegionId() *string 
  SetRelationshipTypesShrink(v string) *EnableCollectionGraphRAGShrinkRequest
  GetRelationshipTypesShrink() *string 
}

type EnableCollectionGraphRAGShrinkRequest struct {
  // The name of the document collection.
  // 
  // > The document collection is created by calling the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation. You can call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to view the created document collections.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // document
  Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
  // The instance ID.
  // 
  // > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including the instance IDs.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // gp-xxxxxxxxx
  DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
  // The list of entity types.
  // 
  // > This parameter is required when knowledge graph construction is enabled.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Location
  EntityTypesShrink *string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty"`
  // The name of the large language model (LLM).
  // 
  // > Valid values:
  // 
  // > - knowledge-extract-standard: Default value.
  // 
  // > - knowledge-extract-mini
  // 
  // > This parameter takes effect only when knowledge graph construction is enabled.
  // 
  // example:
  // 
  // knowledge-extract-standard
  LLMModel *string `json:"LLMModel,omitempty" xml:"LLMModel,omitempty"`
  // The language used to construct the knowledge graph. Valid values:
  // 
  // - Simplified Chinese: Simplified Chinese. Default value.
  // 
  // - English: English.
  // 
  // > This parameter takes effect only when knowledge graph construction is enabled.
  // 
  // example:
  // 
  // Simplified Chinese
  Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
  // The name of the administrator account that has the rds_superuser permission.
  // 
  // > You can call the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) operation to create an account.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // testaccount
  ManagerAccount *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
  // The password of the administrator account.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // testpassword
  ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
  // The namespace. Default value: public.
  // 
  // > You can call the CreateNamespace operation to create a namespace and call the ListNamespaces operation to view the namespace list.
  // 
  // example:
  // 
  // mynamespace
  Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
  // The password for the namespace.
  // 
  // > This value is specified when you call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // testpassword
  NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID.
  // 
  // > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-beijing
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The list of relationship edge types.
  // 
  // > This parameter is required when knowledge graph construction is enabled.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Occurred
  RelationshipTypesShrink *string `json:"RelationshipTypes,omitempty" xml:"RelationshipTypes,omitempty"`
}

func (s EnableCollectionGraphRAGShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCollectionGraphRAGShrinkRequest) GoString() string {
  return s.String()
}

func (s *EnableCollectionGraphRAGShrinkRequest) GetCollection() *string  {
  return s.Collection
}

func (s *EnableCollectionGraphRAGShrinkRequest) GetDBInstanceId() *string  {
  return s.DBInstanceId
}

func (s *EnableCollectionGraphRAGShrinkRequest) GetEntityTypesShrink() *string  {
  return s.EntityTypesShrink
}

func (s *EnableCollectionGraphRAGShrinkRequest) GetLLMModel() *string  {
  return s.LLMModel
}

func (s *EnableCollectionGraphRAGShrinkRequest) GetLanguage() *string  {
  return s.Language
}

func (s *EnableCollectionGraphRAGShrinkRequest) GetManagerAccount() *string  {
  return s.ManagerAccount
}

func (s *EnableCollectionGraphRAGShrinkRequest) GetManagerAccountPassword() *string  {
  return s.ManagerAccountPassword
}

func (s *EnableCollectionGraphRAGShrinkRequest) GetNamespace() *string  {
  return s.Namespace
}

func (s *EnableCollectionGraphRAGShrinkRequest) GetNamespacePassword() *string  {
  return s.NamespacePassword
}

func (s *EnableCollectionGraphRAGShrinkRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableCollectionGraphRAGShrinkRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableCollectionGraphRAGShrinkRequest) GetRelationshipTypesShrink() *string  {
  return s.RelationshipTypesShrink
}

func (s *EnableCollectionGraphRAGShrinkRequest) SetCollection(v string) *EnableCollectionGraphRAGShrinkRequest {
  s.Collection = &v
  return s
}

func (s *EnableCollectionGraphRAGShrinkRequest) SetDBInstanceId(v string) *EnableCollectionGraphRAGShrinkRequest {
  s.DBInstanceId = &v
  return s
}

func (s *EnableCollectionGraphRAGShrinkRequest) SetEntityTypesShrink(v string) *EnableCollectionGraphRAGShrinkRequest {
  s.EntityTypesShrink = &v
  return s
}

func (s *EnableCollectionGraphRAGShrinkRequest) SetLLMModel(v string) *EnableCollectionGraphRAGShrinkRequest {
  s.LLMModel = &v
  return s
}

func (s *EnableCollectionGraphRAGShrinkRequest) SetLanguage(v string) *EnableCollectionGraphRAGShrinkRequest {
  s.Language = &v
  return s
}

func (s *EnableCollectionGraphRAGShrinkRequest) SetManagerAccount(v string) *EnableCollectionGraphRAGShrinkRequest {
  s.ManagerAccount = &v
  return s
}

func (s *EnableCollectionGraphRAGShrinkRequest) SetManagerAccountPassword(v string) *EnableCollectionGraphRAGShrinkRequest {
  s.ManagerAccountPassword = &v
  return s
}

func (s *EnableCollectionGraphRAGShrinkRequest) SetNamespace(v string) *EnableCollectionGraphRAGShrinkRequest {
  s.Namespace = &v
  return s
}

func (s *EnableCollectionGraphRAGShrinkRequest) SetNamespacePassword(v string) *EnableCollectionGraphRAGShrinkRequest {
  s.NamespacePassword = &v
  return s
}

func (s *EnableCollectionGraphRAGShrinkRequest) SetOwnerId(v int64) *EnableCollectionGraphRAGShrinkRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableCollectionGraphRAGShrinkRequest) SetRegionId(v string) *EnableCollectionGraphRAGShrinkRequest {
  s.RegionId = &v
  return s
}

func (s *EnableCollectionGraphRAGShrinkRequest) SetRelationshipTypesShrink(v string) *EnableCollectionGraphRAGShrinkRequest {
  s.RelationshipTypesShrink = &v
  return s
}

func (s *EnableCollectionGraphRAGShrinkRequest) Validate() error {
  return dara.Validate(s)
}

