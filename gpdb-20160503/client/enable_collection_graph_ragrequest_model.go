// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCollectionGraphRAGRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCollection(v string) *EnableCollectionGraphRAGRequest
  GetCollection() *string 
  SetDBInstanceId(v string) *EnableCollectionGraphRAGRequest
  GetDBInstanceId() *string 
  SetEntityTypes(v []*string) *EnableCollectionGraphRAGRequest
  GetEntityTypes() []*string 
  SetLLMModel(v string) *EnableCollectionGraphRAGRequest
  GetLLMModel() *string 
  SetLanguage(v string) *EnableCollectionGraphRAGRequest
  GetLanguage() *string 
  SetManagerAccount(v string) *EnableCollectionGraphRAGRequest
  GetManagerAccount() *string 
  SetManagerAccountPassword(v string) *EnableCollectionGraphRAGRequest
  GetManagerAccountPassword() *string 
  SetNamespace(v string) *EnableCollectionGraphRAGRequest
  GetNamespace() *string 
  SetNamespacePassword(v string) *EnableCollectionGraphRAGRequest
  GetNamespacePassword() *string 
  SetOwnerId(v int64) *EnableCollectionGraphRAGRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableCollectionGraphRAGRequest
  GetRegionId() *string 
  SetRelationshipTypes(v []*string) *EnableCollectionGraphRAGRequest
  GetRelationshipTypes() []*string 
}

type EnableCollectionGraphRAGRequest struct {
  // The name of the document collection.
  // 
  // > You can call the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation to create a document collection and call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to query a list of document collections.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // document
  Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
  // The cluster ID.
  // 
  // > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // gp-xxxxxxxxx
  DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
  // The list of entity types.
  // 
  // > If the knowledge graph construction is enabled, this parameter is required.
  // 
  // This parameter is required.
  EntityTypes []*string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty" type:"Repeated"`
  // The name of the LLM model.
  // 
  // > Valid values:
  // 
  // 	- knowledge-extract-standard: the default value.
  // 
  // 	- knowledge-extract-mini
  // 
  // > This parameter takes effect only when the knowledge graph construction is enabled.
  // 
  // example:
  // 
  // knowledge-extract-standard
  LLMModel *string `json:"LLMModel,omitempty" xml:"LLMModel,omitempty"`
  // The language used to build the knowledge graph. Valid values:
  // 
  // 	- Simplified Chinese. The default value.
  // 
  // 	- English.
  // 
  // > This parameter takes effect only when the knowledge graph construction is enabled.
  // 
  // example:
  // 
  // Simplified Chinese
  Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
  // The name of the privileged database account that has the rds_superuser permission.
  // 
  // > You can call the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) operation to create an account.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // testaccount
  ManagerAccount *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
  // The password for the privileged database account.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // testpassword
  ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
  // The name of the namespace. Default value: public.
  // 
  // > You can call the CreateNamespace operation to create a namespace and call the ListNamespaces operation to query a list of namespaces.
  // 
  // example:
  // 
  // mynamespace
  Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
  // The password of the namespace.
  // 
  // > The value of this parameter is specified by [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // testpassword
  NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID of the cluster.
  // 
  // > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-beijing
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The list of relationship edge types.
  // 
  // > If the knowledge graph construction is enabled, this parameter is required.
  // 
  // This parameter is required.
  RelationshipTypes []*string `json:"RelationshipTypes,omitempty" xml:"RelationshipTypes,omitempty" type:"Repeated"`
}

func (s EnableCollectionGraphRAGRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCollectionGraphRAGRequest) GoString() string {
  return s.String()
}

func (s *EnableCollectionGraphRAGRequest) GetCollection() *string  {
  return s.Collection
}

func (s *EnableCollectionGraphRAGRequest) GetDBInstanceId() *string  {
  return s.DBInstanceId
}

func (s *EnableCollectionGraphRAGRequest) GetEntityTypes() []*string  {
  return s.EntityTypes
}

func (s *EnableCollectionGraphRAGRequest) GetLLMModel() *string  {
  return s.LLMModel
}

func (s *EnableCollectionGraphRAGRequest) GetLanguage() *string  {
  return s.Language
}

func (s *EnableCollectionGraphRAGRequest) GetManagerAccount() *string  {
  return s.ManagerAccount
}

func (s *EnableCollectionGraphRAGRequest) GetManagerAccountPassword() *string  {
  return s.ManagerAccountPassword
}

func (s *EnableCollectionGraphRAGRequest) GetNamespace() *string  {
  return s.Namespace
}

func (s *EnableCollectionGraphRAGRequest) GetNamespacePassword() *string  {
  return s.NamespacePassword
}

func (s *EnableCollectionGraphRAGRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableCollectionGraphRAGRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableCollectionGraphRAGRequest) GetRelationshipTypes() []*string  {
  return s.RelationshipTypes
}

func (s *EnableCollectionGraphRAGRequest) SetCollection(v string) *EnableCollectionGraphRAGRequest {
  s.Collection = &v
  return s
}

func (s *EnableCollectionGraphRAGRequest) SetDBInstanceId(v string) *EnableCollectionGraphRAGRequest {
  s.DBInstanceId = &v
  return s
}

func (s *EnableCollectionGraphRAGRequest) SetEntityTypes(v []*string) *EnableCollectionGraphRAGRequest {
  s.EntityTypes = v
  return s
}

func (s *EnableCollectionGraphRAGRequest) SetLLMModel(v string) *EnableCollectionGraphRAGRequest {
  s.LLMModel = &v
  return s
}

func (s *EnableCollectionGraphRAGRequest) SetLanguage(v string) *EnableCollectionGraphRAGRequest {
  s.Language = &v
  return s
}

func (s *EnableCollectionGraphRAGRequest) SetManagerAccount(v string) *EnableCollectionGraphRAGRequest {
  s.ManagerAccount = &v
  return s
}

func (s *EnableCollectionGraphRAGRequest) SetManagerAccountPassword(v string) *EnableCollectionGraphRAGRequest {
  s.ManagerAccountPassword = &v
  return s
}

func (s *EnableCollectionGraphRAGRequest) SetNamespace(v string) *EnableCollectionGraphRAGRequest {
  s.Namespace = &v
  return s
}

func (s *EnableCollectionGraphRAGRequest) SetNamespacePassword(v string) *EnableCollectionGraphRAGRequest {
  s.NamespacePassword = &v
  return s
}

func (s *EnableCollectionGraphRAGRequest) SetOwnerId(v int64) *EnableCollectionGraphRAGRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableCollectionGraphRAGRequest) SetRegionId(v string) *EnableCollectionGraphRAGRequest {
  s.RegionId = &v
  return s
}

func (s *EnableCollectionGraphRAGRequest) SetRelationshipTypes(v []*string) *EnableCollectionGraphRAGRequest {
  s.RelationshipTypes = v
  return s
}

func (s *EnableCollectionGraphRAGRequest) Validate() error {
  return dara.Validate(s)
}

