// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEngineSearchRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppId(v string) *EngineSearchRequest
  GetAppId() *string 
  SetGrey(v bool) *EngineSearchRequest
  GetGrey() *bool 
  SetQuery(v *EngineSearchRequestQuery) *EngineSearchRequest
  GetQuery() *EngineSearchRequestQuery 
  SetRecall(v *EngineSearchRequestRecall) *EngineSearchRequest
  GetRecall() *EngineSearchRequestRecall 
  SetSessionId(v string) *EngineSearchRequest
  GetSessionId() *string 
  SetStrategyId(v string) *EngineSearchRequest
  GetStrategyId() *string 
  SetUser(v *EngineSearchRequestUser) *EngineSearchRequest
  GetUser() *EngineSearchRequestUser 
  SetVersion(v string) *EngineSearchRequest
  GetVersion() *string 
}

type EngineSearchRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 507218
  AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
  // example:
  // 
  // false
  Grey *bool `json:"grey,omitempty" xml:"grey,omitempty"`
  // This parameter is required.
  Query *EngineSearchRequestQuery `json:"query,omitempty" xml:"query,omitempty" type:"Struct"`
  Recall *EngineSearchRequestRecall `json:"recall,omitempty" xml:"recall,omitempty" type:"Struct"`
  // example:
  // 
  // 2e95ef4fbc28437db5008a910bd392a4
  SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
  // example:
  // 
  // 507218
  StrategyId *string `json:"strategyId,omitempty" xml:"strategyId,omitempty"`
  User *EngineSearchRequestUser `json:"user,omitempty" xml:"user,omitempty" type:"Struct"`
  // example:
  // 
  // 默认1.0
  Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s EngineSearchRequest) String() string {
  return dara.Prettify(s)
}

func (s EngineSearchRequest) GoString() string {
  return s.String()
}

func (s *EngineSearchRequest) GetAppId() *string  {
  return s.AppId
}

func (s *EngineSearchRequest) GetGrey() *bool  {
  return s.Grey
}

func (s *EngineSearchRequest) GetQuery() *EngineSearchRequestQuery  {
  return s.Query
}

func (s *EngineSearchRequest) GetRecall() *EngineSearchRequestRecall  {
  return s.Recall
}

func (s *EngineSearchRequest) GetSessionId() *string  {
  return s.SessionId
}

func (s *EngineSearchRequest) GetStrategyId() *string  {
  return s.StrategyId
}

func (s *EngineSearchRequest) GetUser() *EngineSearchRequestUser  {
  return s.User
}

func (s *EngineSearchRequest) GetVersion() *string  {
  return s.Version
}

func (s *EngineSearchRequest) SetAppId(v string) *EngineSearchRequest {
  s.AppId = &v
  return s
}

func (s *EngineSearchRequest) SetGrey(v bool) *EngineSearchRequest {
  s.Grey = &v
  return s
}

func (s *EngineSearchRequest) SetQuery(v *EngineSearchRequestQuery) *EngineSearchRequest {
  s.Query = v
  return s
}

func (s *EngineSearchRequest) SetRecall(v *EngineSearchRequestRecall) *EngineSearchRequest {
  s.Recall = v
  return s
}

func (s *EngineSearchRequest) SetSessionId(v string) *EngineSearchRequest {
  s.SessionId = &v
  return s
}

func (s *EngineSearchRequest) SetStrategyId(v string) *EngineSearchRequest {
  s.StrategyId = &v
  return s
}

func (s *EngineSearchRequest) SetUser(v *EngineSearchRequestUser) *EngineSearchRequest {
  s.User = v
  return s
}

func (s *EngineSearchRequest) SetVersion(v string) *EngineSearchRequest {
  s.Version = &v
  return s
}

func (s *EngineSearchRequest) Validate() error {
  if s.Query != nil {
    if err := s.Query.Validate(); err != nil {
      return err
    }
  }
  if s.Recall != nil {
    if err := s.Recall.Validate(); err != nil {
      return err
    }
  }
  if s.User != nil {
    if err := s.User.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EngineSearchRequestQuery struct {
  // example:
  // 
  // ["24234111","12423525"]
  ExcludeIds []*string `json:"excludeIds,omitempty" xml:"excludeIds,omitempty" type:"Repeated"`
  // example:
  // 
  // ["https://paperreview.oss-cn-hangzhou.aliyuncs.com/59dd424f-97ed-4855-942e-c961f1f5b67e.jpeg"]
  ImageUrls []*string `json:"imageUrls,omitempty" xml:"imageUrls,omitempty" type:"Repeated"`
  // example:
  // 
  // 1
  PageNo *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // example:
  // 
  // 10
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // example:
  // 
  // ["梅花图片"]
  Texts []*string `json:"texts,omitempty" xml:"texts,omitempty" type:"Repeated"`
}

func (s EngineSearchRequestQuery) String() string {
  return dara.Prettify(s)
}

func (s EngineSearchRequestQuery) GoString() string {
  return s.String()
}

func (s *EngineSearchRequestQuery) GetExcludeIds() []*string  {
  return s.ExcludeIds
}

func (s *EngineSearchRequestQuery) GetImageUrls() []*string  {
  return s.ImageUrls
}

func (s *EngineSearchRequestQuery) GetPageNo() *int32  {
  return s.PageNo
}

func (s *EngineSearchRequestQuery) GetPageSize() *int32  {
  return s.PageSize
}

func (s *EngineSearchRequestQuery) GetTexts() []*string  {
  return s.Texts
}

func (s *EngineSearchRequestQuery) SetExcludeIds(v []*string) *EngineSearchRequestQuery {
  s.ExcludeIds = v
  return s
}

func (s *EngineSearchRequestQuery) SetImageUrls(v []*string) *EngineSearchRequestQuery {
  s.ImageUrls = v
  return s
}

func (s *EngineSearchRequestQuery) SetPageNo(v int32) *EngineSearchRequestQuery {
  s.PageNo = &v
  return s
}

func (s *EngineSearchRequestQuery) SetPageSize(v int32) *EngineSearchRequestQuery {
  s.PageSize = &v
  return s
}

func (s *EngineSearchRequestQuery) SetTexts(v []*string) *EngineSearchRequestQuery {
  s.Texts = v
  return s
}

func (s *EngineSearchRequestQuery) Validate() error {
  return dara.Validate(s)
}

type EngineSearchRequestRecall struct {
  // example:
  // 
  // false
  CloseRecallAsr *bool `json:"closeRecallAsr,omitempty" xml:"closeRecallAsr,omitempty"`
  // example:
  // 
  // true
  NeedMergeSegments *bool `json:"needMergeSegments,omitempty" xml:"needMergeSegments,omitempty"`
}

func (s EngineSearchRequestRecall) String() string {
  return dara.Prettify(s)
}

func (s EngineSearchRequestRecall) GoString() string {
  return s.String()
}

func (s *EngineSearchRequestRecall) GetCloseRecallAsr() *bool  {
  return s.CloseRecallAsr
}

func (s *EngineSearchRequestRecall) GetNeedMergeSegments() *bool  {
  return s.NeedMergeSegments
}

func (s *EngineSearchRequestRecall) SetCloseRecallAsr(v bool) *EngineSearchRequestRecall {
  s.CloseRecallAsr = &v
  return s
}

func (s *EngineSearchRequestRecall) SetNeedMergeSegments(v bool) *EngineSearchRequestRecall {
  s.NeedMergeSegments = &v
  return s
}

func (s *EngineSearchRequestRecall) Validate() error {
  return dara.Validate(s)
}

type EngineSearchRequestUser struct {
  // example:
  // 
  // "123456"
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s EngineSearchRequestUser) String() string {
  return dara.Prettify(s)
}

func (s EngineSearchRequestUser) GoString() string {
  return s.String()
}

func (s *EngineSearchRequestUser) GetUserId() *string  {
  return s.UserId
}

func (s *EngineSearchRequestUser) SetUserId(v string) *EngineSearchRequestUser {
  s.UserId = &v
  return s
}

func (s *EngineSearchRequestUser) Validate() error {
  return dara.Validate(s)
}

