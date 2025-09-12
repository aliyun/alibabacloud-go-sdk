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
  // This parameter is required.
  // 
  // example:
  // 
  // document
  Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // gp-xxxxxxxxx
  DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
  // This parameter is required.
  EntityTypes []*string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty" type:"Repeated"`
  // example:
  // 
  // knowledge-extract-standard
  LLMModel *string `json:"LLMModel,omitempty" xml:"LLMModel,omitempty"`
  // example:
  // 
  // Simplified Chinese
  Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // testaccount
  ManagerAccount *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // testpassword
  ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
  // example:
  // 
  // mynamespace
  Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // testpassword
  NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // cn-beijing
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

