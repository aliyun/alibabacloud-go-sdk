// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateTicketRequest interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *EvaluateTicketRequest
  GetContent() *string 
  SetScore(v string) *EvaluateTicketRequest
  GetScore() *string 
  SetSolved(v bool) *EvaluateTicketRequest
  GetSolved() *bool 
  SetTicketId(v string) *EvaluateTicketRequest
  GetTicketId() *string 
  SetUid(v string) *EvaluateTicketRequest
  GetUid() *string 
}

type EvaluateTicketRequest struct {
  // Comment
  // 
  // example:
  // 
  // The engineer solved my issue.
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
  // Rating star 1-5 stars
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
  // Whether to resolve
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // true
  Solved *bool `json:"Solved,omitempty" xml:"Solved,omitempty"`
  // The ID of the ticket.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 001ET1BU1P
  TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
  // UID
  // 
  // example:
  // 
  // 1902070573958003
  Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s EvaluateTicketRequest) String() string {
  return dara.Prettify(s)
}

func (s EvaluateTicketRequest) GoString() string {
  return s.String()
}

func (s *EvaluateTicketRequest) GetContent() *string  {
  return s.Content
}

func (s *EvaluateTicketRequest) GetScore() *string  {
  return s.Score
}

func (s *EvaluateTicketRequest) GetSolved() *bool  {
  return s.Solved
}

func (s *EvaluateTicketRequest) GetTicketId() *string  {
  return s.TicketId
}

func (s *EvaluateTicketRequest) GetUid() *string  {
  return s.Uid
}

func (s *EvaluateTicketRequest) SetContent(v string) *EvaluateTicketRequest {
  s.Content = &v
  return s
}

func (s *EvaluateTicketRequest) SetScore(v string) *EvaluateTicketRequest {
  s.Score = &v
  return s
}

func (s *EvaluateTicketRequest) SetSolved(v bool) *EvaluateTicketRequest {
  s.Solved = &v
  return s
}

func (s *EvaluateTicketRequest) SetTicketId(v string) *EvaluateTicketRequest {
  s.TicketId = &v
  return s
}

func (s *EvaluateTicketRequest) SetUid(v string) *EvaluateTicketRequest {
  s.Uid = &v
  return s
}

func (s *EvaluateTicketRequest) Validate() error {
  return dara.Validate(s)
}

