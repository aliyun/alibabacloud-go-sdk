// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditBiddingDocRequest interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *EditBiddingDocRequest
  GetContent() *string 
  SetContentFormat(v string) *EditBiddingDocRequest
  GetContentFormat() *string 
  SetContentType(v string) *EditBiddingDocRequest
  GetContentType() *string 
  SetTaskId(v string) *EditBiddingDocRequest
  GetTaskId() *string 
  SetWorkspaceId(v string) *EditBiddingDocRequest
  GetWorkspaceId() *string 
}

type EditBiddingDocRequest struct {
  // The text content.
  // 
  // example:
  // 
  // 标书内容
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
  // The format.
  // 
  // example:
  // 
  // html
  // 
  // markdown
  ContentFormat *string `json:"ContentFormat,omitempty" xml:"ContentFormat,omitempty"`
  // The content type.
  // 
  // example:
  // 
  // outline
  // 
  // bidding
  ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
  // The task ID.
  // 
  // example:
  // 
  // 0dbf1055f8a2475d99904c3b76a0ffba
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
  // [The workspace ID.](https://help.aliyun.com/document_detail/2782167.html)
  // 
  // example:
  // 
  // llm-xx
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s EditBiddingDocRequest) String() string {
  return dara.Prettify(s)
}

func (s EditBiddingDocRequest) GoString() string {
  return s.String()
}

func (s *EditBiddingDocRequest) GetContent() *string  {
  return s.Content
}

func (s *EditBiddingDocRequest) GetContentFormat() *string  {
  return s.ContentFormat
}

func (s *EditBiddingDocRequest) GetContentType() *string  {
  return s.ContentType
}

func (s *EditBiddingDocRequest) GetTaskId() *string  {
  return s.TaskId
}

func (s *EditBiddingDocRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *EditBiddingDocRequest) SetContent(v string) *EditBiddingDocRequest {
  s.Content = &v
  return s
}

func (s *EditBiddingDocRequest) SetContentFormat(v string) *EditBiddingDocRequest {
  s.ContentFormat = &v
  return s
}

func (s *EditBiddingDocRequest) SetContentType(v string) *EditBiddingDocRequest {
  s.ContentType = &v
  return s
}

func (s *EditBiddingDocRequest) SetTaskId(v string) *EditBiddingDocRequest {
  s.TaskId = &v
  return s
}

func (s *EditBiddingDocRequest) SetWorkspaceId(v string) *EditBiddingDocRequest {
  s.WorkspaceId = &v
  return s
}

func (s *EditBiddingDocRequest) Validate() error {
  return dara.Validate(s)
}

