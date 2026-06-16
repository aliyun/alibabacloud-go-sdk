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
  SetSessionId(v string) *EngineSearchRequest
  GetSessionId() *string 
  SetUser(v *EngineSearchRequestUser) *EngineSearchRequest
  GetUser() *EngineSearchRequestUser 
}

type EngineSearchRequest struct {
  // The unique ID of the application.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 2048962366415007746
  AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
  // Specifies whether to access the draft version.
  // 
  // example:
  // 
  // false
  Grey *bool `json:"grey,omitempty" xml:"grey,omitempty"`
  // The query condition object.
  // 
  // This parameter is required.
  Query *EngineSearchRequestQuery `json:"query,omitempty" xml:"query,omitempty" type:"Struct"`
  // This parameter does not need to be specified.
  // 
  // example:
  // 
  // 2e95ef4fbc28437db5008a910bd392a4
  SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
  // The user information object, used for subsequent user-perspective analysis.
  User *EngineSearchRequestUser `json:"user,omitempty" xml:"user,omitempty" type:"Struct"`
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

func (s *EngineSearchRequest) GetSessionId() *string  {
  return s.SessionId
}

func (s *EngineSearchRequest) GetUser() *EngineSearchRequestUser  {
  return s.User
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

func (s *EngineSearchRequest) SetSessionId(v string) *EngineSearchRequest {
  s.SessionId = &v
  return s
}

func (s *EngineSearchRequest) SetUser(v *EngineSearchRequestUser) *EngineSearchRequest {
  s.User = v
  return s
}

func (s *EngineSearchRequest) Validate() error {
  if s.Query != nil {
    if err := s.Query.Validate(); err != nil {
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
  // The list of primary key IDs to exclude.<br>• Purpose: filters out previously viewed history records.
  // 
  // example:
  // 
  // ["24234111","12423525"]
  ExcludeIds []*string `json:"excludeIds,omitempty" xml:"excludeIds,omitempty" type:"Repeated"`
  // The image query list.<br>• Only one image URL is supported. The maximum size of a single image is 10 MB. Supported formats: JPG, PNG, WEBP, and JPEG.
  // 
  // example:
  // 
  // ["https://paperreview.oss-cn-hangzhou.aliyuncs.com/59dd424f-97ed-4855-942e-c961f1f5b67e.jpeg"]
  ImageUrls []*string `json:"imageUrls,omitempty" xml:"imageUrls,omitempty" type:"Repeated"`
  // The page number, starting from 1.<br>• Default value: `1`.
  // 
  // example:
  // 
  // 1
  PageNo *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // The number of results returned per page.
  // 
  // example:
  // 
  // 10
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // The text query list.<br>• Only one text string is supported. The maximum length is 256 characters.
  // 
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

type EngineSearchRequestUser struct {
  // The unique ID of the user.
  // 
  // example:
  // 
  // asdfgnoevnor
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

