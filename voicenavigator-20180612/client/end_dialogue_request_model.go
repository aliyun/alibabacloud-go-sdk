// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndDialogueRequest interface {
  dara.Model
  String() string
  GoString() string
  SetConversationId(v string) *EndDialogueRequest
  GetConversationId() *string 
  SetHangUpParams(v string) *EndDialogueRequest
  GetHangUpParams() *string 
  SetInstanceId(v string) *EndDialogueRequest
  GetInstanceId() *string 
  SetInstanceOwnerId(v int64) *EndDialogueRequest
  GetInstanceOwnerId() *int64 
}

type EndDialogueRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 8fb819b5-d032-48a9-ae5e-cff041b83596
  ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
  // example:
  // 
  // {\\"duration\\":40,\\"endTime\\":1645082505345,\\"hangUpDirection\\":\\"ivr\\",\\"hasLastPlaybackCompleted\\":true,\\"startTime\\":1645082505305}
  HangUpParams *string `json:"HangUpParams,omitempty" xml:"HangUpParams,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // e48e45dd-e47a-4744-a063-f08cbebb1c5a
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // example:
  // 
  // 1399572315967217
  InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
}

func (s EndDialogueRequest) String() string {
  return dara.Prettify(s)
}

func (s EndDialogueRequest) GoString() string {
  return s.String()
}

func (s *EndDialogueRequest) GetConversationId() *string  {
  return s.ConversationId
}

func (s *EndDialogueRequest) GetHangUpParams() *string  {
  return s.HangUpParams
}

func (s *EndDialogueRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EndDialogueRequest) GetInstanceOwnerId() *int64  {
  return s.InstanceOwnerId
}

func (s *EndDialogueRequest) SetConversationId(v string) *EndDialogueRequest {
  s.ConversationId = &v
  return s
}

func (s *EndDialogueRequest) SetHangUpParams(v string) *EndDialogueRequest {
  s.HangUpParams = &v
  return s
}

func (s *EndDialogueRequest) SetInstanceId(v string) *EndDialogueRequest {
  s.InstanceId = &v
  return s
}

func (s *EndDialogueRequest) SetInstanceOwnerId(v int64) *EndDialogueRequest {
  s.InstanceOwnerId = &v
  return s
}

func (s *EndDialogueRequest) Validate() error {
  return dara.Validate(s)
}

