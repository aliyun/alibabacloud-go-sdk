// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iSessionCreate interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *SessionCreate
  GetType() *string 
  SetEventId(v string) *SessionCreate
  GetEventId() *string 
  SetMode(v string) *SessionCreate
  GetMode() *string 
  SetLanguagePair(v []*string) *SessionCreate
  GetLanguagePair() []*string 
  SetSession(v *SessionCreateSession) *SessionCreate
  GetSession() *SessionCreateSession 
}

type SessionCreate struct {
  // This parameter is required.
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // 可省略；提供时必须为非空、非纯空白字符串，用于关联响应。
  // 
  // example:
  // 
  // evt_create_1
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  // 与 session.mode 二选一；同时提供必须一致。均省略时后端选择 single，网关不补默认值。
  Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
  // 双向条件必填，与 session.language_pair 二选一；同时提供时内容及顺序必须一致。仅支持 zh 与 en/ja/ko 组合，顺序可反转，由后端校验。
  // 
  // example:
  // 
  // ["zh","en"]
  LanguagePair []*string `json:"language_pair,omitempty" xml:"language_pair,omitempty" type:"Repeated"`
  // 单向可省略或指定 mode=single；双向须指定 bidirectional 模式及语言对。创建后模式与语言对不可变，音频和输出配置通过 SessionUpdate 下发。
  // 
  // example:
  // 
  // {"mode":"bidirectional","language_pair":["zh","en"]}
  Session *SessionCreateSession `json:"session,omitempty" xml:"session,omitempty" type:"Struct"`
}

func (s SessionCreate) String() string {
  return dara.Prettify(s)
}

func (s SessionCreate) GoString() string {
  return s.String()
}

func (s *SessionCreate) GetType() *string  {
  return s.Type
}

func (s *SessionCreate) GetEventId() *string  {
  return s.EventId
}

func (s *SessionCreate) GetMode() *string  {
  return s.Mode
}

func (s *SessionCreate) GetLanguagePair() []*string  {
  return s.LanguagePair
}

func (s *SessionCreate) GetSession() *SessionCreateSession  {
  return s.Session
}

func (s *SessionCreate) SetType(v string) *SessionCreate {
  s.Type = &v
  return s
}

func (s *SessionCreate) SetEventId(v string) *SessionCreate {
  s.EventId = &v
  return s
}

func (s *SessionCreate) SetMode(v string) *SessionCreate {
  s.Mode = &v
  return s
}

func (s *SessionCreate) SetLanguagePair(v []*string) *SessionCreate {
  s.LanguagePair = v
  return s
}

func (s *SessionCreate) SetSession(v *SessionCreateSession) *SessionCreate {
  s.Session = v
  return s
}

func (s *SessionCreate) Validate() error {
  if s.Session != nil {
    if err := s.Session.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionCreateSession struct {
  Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
  // 双向须提供本字段或根节点 language_pair；须包含 zh 和 en/ja/ko 之一。
  LanguagePair []*string `json:"language_pair,omitempty" xml:"language_pair,omitempty" type:"Repeated"`
}

func (s SessionCreateSession) String() string {
  return dara.Prettify(s)
}

func (s SessionCreateSession) GoString() string {
  return s.String()
}

func (s *SessionCreateSession) GetMode() *string  {
  return s.Mode
}

func (s *SessionCreateSession) GetLanguagePair() []*string  {
  return s.LanguagePair
}

func (s *SessionCreateSession) SetMode(v string) *SessionCreateSession {
  s.Mode = &v
  return s
}

func (s *SessionCreateSession) SetLanguagePair(v []*string) *SessionCreateSession {
  s.LanguagePair = v
  return s
}

func (s *SessionCreateSession) Validate() error {
  return dara.Validate(s)
}

